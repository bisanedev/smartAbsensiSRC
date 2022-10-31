import React from 'react';
import {    
  Button,  
  StyleSheet,
  View,
  Text, 
  RefreshControl,
  ScrollView,  
} from 'react-native';
import AsyncStorage from '@react-native-async-storage/async-storage';
import MyStatusBar from '../../components/statusBar';
import moment from 'moment';
import axios from 'axios';
import Toast, {DURATION} from 'react-native-easy-toast';
import Header from './editheader';
import Body from '../absensi/body';
import update from 'immutability-helper';

export default class ArsipEditScreen extends React.Component {  

    constructor(props) {
        super(props);     
       
        this.state = {siswa:[],refreshing: true,ipaddress: '',tombol:''};        
       
        
        this.dataEdit = this.props.route.params.dataEdit;         
        //---
    }

    componentDidMount(){
        AsyncStorage.getItem('ipaddress', (error, result) => {
            if (result) {            
                this.setState({
                    ipaddress: result
                });
                this._SiswaData();
            }          
        });     
    }
  
    render() {        
        // tombol
        let tombol;
        if(this.state.tombol === "yes")
        {tombol = <Button title="Update Absen" color="#757575" onPress={this._AbsenData} />}
        else
        {tombol = <Button title="Simpan Absen" color="#757575" onPress={this._AbsenData} />}        
        //---- end
        return (        
        <View style={styles.container}>
          <MyStatusBar backgroundColor="#fff" barStyle="dark-content" />
          <Header kembali={this._goBack} piket={this._handlePiket} disablepiket={this.piket}/>
          <View style={styles.body}>
          <ScrollView refreshControl={
            <RefreshControl refreshing={this.state.refreshing} onRefresh={this.onRefresh.bind(this)} />
          }>                              
          <View style={styles.card}>                               
            <View style={styles.konten}>
                <Text style={styles.txtnama}>{this.dataEdit.mapel_guru}</Text>                                        
                <Text style={styles.txt}>Mata Pelajaran {this.dataEdit.mapel}</Text>
                <Text style={styles.txt}>Jam {moment(this.dataEdit.mapel_start).format('HH:mm')} ~ {moment(this.dataEdit.mapel_end).format('HH:mm')} , Mata Pelajaran Ke {this.dataEdit.jam}</Text>                
                <Text style={styles.txttgl}>Tanggal {moment(this.dataEdit.tanggal).format('DD/MM/YYYY')} [{moment(this.dataEdit.tanggal).format('HH:mm')}~{moment(this.dataEdit.keluar).format('HH:mm')}]</Text>                
            </View>                       
          </View>
          <View style={styles.headerinfo}>   
              <View style={styles.colum}>
                <Text style={{color:'#404040',fontWeight:'bold'}}> Kelas {this.dataEdit.kelas}</Text>
              </View>
              <View style={styles.colum}>
                <Text style={{color:'#404040',fontWeight:'bold'}}>{this.state.siswa.length} Siswa</Text>
              </View>                           
          </View>                     
          <View style={styles.kontensiswa}>              
          { this.state.siswa.sort((a, b) => a.urut > b.urut).map((item,i) => (
              <Body data={item} key={i} absen={this._absen}/>
          ))}                          
          </View>
          <View style={this.state.refreshing ? styles.hide : styles.menubawah}>                    
          <View style={styles.infoabsen}>
            <View style={styles.columbawah}>
            <Text style={styles.columbawahTxt}>Hadir</Text>
            <Text style={styles.jumlah}>{this._countJumlah('h')}</Text>
            </View>
            <View style={styles.columbawah}>
            <Text style={styles.columbawahTxt}>Sakit</Text>
            <Text style={styles.jumlah}>{this._countJumlah('s')}</Text>
            </View>
            <View style={styles.columbawah}>
            <Text style={styles.columbawahTxt}>Izin</Text>
            <Text style={styles.jumlah}>{this._countJumlah('i')}</Text>
            </View>
            <View style={styles.columbawah}>
            <Text style={styles.columbawahTxt}>Terlambat</Text>
            <Text style={styles.jumlah}>{this._countJumlah('t')}</Text>
            </View>
            <View style={styles.columbawah}>              
            <Text style={styles.columbawahTxt}>Alpa</Text>
            <Text style={styles.jumlah}>{this._countJumlah('a')}</Text>
            </View>            
          </View>                   
          <View style={styles.tombol}>
            {tombol}
          </View>
          </View>          
          </ScrollView>  
          </View>          
          <Toast ref="toast" defaultCloseDelay={1000} position='top' positionValue={200} opacity={0.8}/>        
        </View>
      );
    }          
    //function
    _goBack = () => {            
      this.props.navigation.goBack();
    };

    _absen = (no,value,note) => {
      var data = this.state.siswa;     
      for (var key in data) {
          if (data.hasOwnProperty(key)) {  
            if(data[key].siswa_id == no){
              var obj = update(this.state.siswa, {
                [key]:{kode:{$set:value},kode_note:{$set:note}}
              });
              this.setState({siswa: obj});
            }           
        }
      } 
    }

    _countJumlah = (kode) => {
      var data = this.state.siswa; 
      return data.filter(item => item.kode === kode ).length      
    }

    onRefresh = async () => { 
      await this.setState({ siswa:[]});      
      //Call the Service to get the latest data
      this._SiswaData();
    };

    _SiswaData = async ()  => {                 
      await axios.get("http://"+this.state.ipaddress+"/api/absensi/getsiswa?kelas="+this.dataEdit.kelas_id+"&tanggal="+moment(this.dataEdit.tanggal).format('YYYY-MM-DDTHH:mm')+"&jam="+this.dataEdit.jam)
         .then(response => {                                                     
           this.setState({ siswa: response.data.siswa });            
           this.setState({ refreshing: false });           
           this.setState({ tombol: response.data.updated });  
         }).catch(error => {           
          
          if(error.message === "Network Error"){            
              this.refs.toast.show('Maaf Jaringan Tidak Tersedia !');
          }
          if(error.response.status == 401){
              this.refs.toast.show(error.response.data.message);               
          }
          
        });      
      // ---
    };

   
    _AbsenData = async ()  => {                  
      await axios.post("http://"+this.state.ipaddress+"/api/absensi/mobile/store",
      {
          tanggal:moment(this.dataEdit.tanggal).format('YYYY-MM-DDTHH:mm'),user_id:Number(this.dataEdit.user_id),kelas_id:Number(this.dataEdit.kelas_id),jam:Number(this.dataEdit.jam),kelas:this.dataEdit.kelas,
          mapel:this.dataEdit.mapel,mapelguru:this.dataEdit.mapel_guru,mapel_start:this.dataEdit.mapel_start,mapel_end:this.dataEdit.mapel_end,mapelnote:this.dataEdit.mapel_note,
          siswa:JSON.stringify(this.state.siswa)
      }).then(response => { 
          if(response.data.status == "success")
            { 
              this.refs.toast.show('Absensi Tersimpan Di Server !');
            }
         }).catch(error => {           
          
          if(error.message === "Network Error"){            
              this.refs.toast.show('Maaf Jaringan Tidak Tersedia !');
          }
          if(error.response.status == 401){
              this.refs.toast.show(error.response.data.message);               
          }
          
        });
      
      // ---
    };          
    // -----    
  }
  
  const styles = StyleSheet.create({
    container: {      
      flex: 1,
    },   
    body:{
      flex:1,      
      backgroundColor:'#f2f1f1',      
    },
    card:{
      backgroundColor:'#f9f9f9',      
      flexDirection:'row', 
      justifyContent:'space-between', 
      alignItems:'center',                                            
    },
    infoabsen:{     
      flex: 1,
      marginTop:7,
      marginBottom:7,
      backgroundColor: '#f9f9f9',
      flexDirection: 'row'
    },    
    menubawah:{
      backgroundColor:'#f9f9f9',
    },   
    columbawah:{
      width:'20%',         
      textAlign:'center',
      alignSelf: 'center',
      alignItems: 'center',
      justifyContent: 'center', 
      height:40,   
    },
    columbawahTxt:{
      color:'#404040',
      fontWeight:'bold',
      fontSize:10,
      textTransform:'uppercase',
    },
    jumlah:{
      fontSize:11,      
      textAlign:'center',        
    },
    pic:{        
        width:80,
        height:120,        
    },
    tombol:{
      marginTop:5,
      marginBottom:20,
      marginLeft:10,
      marginRight:10,
    },
    hide:{
      display: 'none',       
    },
    colum:{
      width:'50%',
      borderTopColor: '#e6e6e6',
      borderTopWidth: 0.5,
      borderBottomColor: '#e6e6e6',
      borderBottomWidth: 0.5,
      borderRightColor:'#e6e6e6',
      borderRightWidth:0.5,          
      textAlign:'center',
      alignSelf: 'center',
      alignItems: 'center',
      justifyContent: 'center', 
      height:40,   
    },
    headerinfo:{            
      flex: 1,
      backgroundColor: '#f9f9f9',
      flexDirection: 'row',
    },
    konten:{              
        flex: 1,        
        alignItems: 'center',        
    },
    kontensiswa:{
      backgroundColor:'#f2f1f1',
      marginBottom:15,
    },
    txt:{
      textTransform:'capitalize',
      marginTop:3,    
      fontSize:11,
      color:'#757575',
    },
    txttgl:{
        marginTop:8,
        marginBottom:10,        
        fontSize:13,
        fontWeight:'700',
        color:'#404040',
    },
    txtnama:{  
        textTransform:'capitalize',      
        fontWeight:'bold',
        fontSize:18,
        marginBottom:5,        
        color:'#404040',
    },
  });