import React from 'react';
import {    
  Button,  
  StyleSheet,
  View,
  Text,
  Alert,  
  Image,
  RefreshControl,
  ScrollView,
  BackHandler,
} from 'react-native';
import AsyncStorage from '@react-native-async-storage/async-storage';
import MyStatusBar from '../../components/statusBar';
import withPreventDoubleClick from '../../components/withPreventDoubleClick';
import moment from 'moment';
import axios from 'axios';
import Toast, {DURATION} from 'react-native-easy-toast';
import Header from './header';
import Body from './body';
import update from 'immutability-helper';

const ButtonPrevent = withPreventDoubleClick(Button);

export default class AbsenScreen extends React.Component {  

    constructor(props) {
        super(props);     
       
        this.state = {siswa:[],refreshing: true,ipaddress: '',piketguru:'',tombol:'',lastupdate:''};      

        AsyncStorage.getItem('ipaddress', (error, result) => {
          if (result) {            
              this.setState({
                  ipaddress: result
              });              
              this.gurupiket = this.props.route.params.piket;
              
              if (this.gurupiket === undefined) {
                this._mapelData();
              }else{
                this._mapelData(this.gurupiket);
              }
          }          
        });
        this._isMounted = false;
        this.data = this.props.route.params.data; 
        this.piket = this.props.route.params.disablepiket;         
        //---
    }

  
    render() {
        const cewek = <Image style={styles.pic} source={{uri: "http://"+this.state.ipaddress+"/assets/images/cewek.png"}} />;
        const cowok = <Image style={styles.pic} source={{uri: "http://"+this.state.ipaddress+"/assets/images/cowok.png"}} />;
        const photo = <Image style={styles.pic} source={{uri: "http://"+this.state.ipaddress+"/storage/guru/"+this.data.Users.foto}} />;
        let profile;
        if(this.data.Users.foto){ profile = photo }
        else if(this.data.Users.jenis === "Perempuan"){profile = cewek}else{profile = cowok}
        // tombol
        let tombol;
        if(this.state.tombol === "yes")
        {tombol = <ButtonPrevent title="Perbarui Absen" color="#757575" onPress={this._AbsenData} />}               
        //---- end
        return (        
        <View style={styles.container}>
          <MyStatusBar backgroundColor="#fff" barStyle="dark-content" />
          <Header kembali={this._goBack} piket={this.siapapiket} disablepiket={this.piket}/>
          <View style={styles.body}>
          <ScrollView refreshControl={
            <RefreshControl refreshing={this.state.refreshing} onRefresh={this.onRefresh.bind(this)} />
          }>                              
          <View style={styles.card}>
            {profile}                    
            <View style={styles.konten}>
                <Text style={styles.txtnama}>{this.data.Users.nama}</Text>                                        
                <Text style={styles.txt}>Mata Pelajaran {this.data.Users.mapel}</Text>
                <Text style={styles.txt}>Hari {this.data.hari} , Jam {this.data.jam_start} ~ {this.data.jam_end}</Text>
                <Text style={styles.txt}>Jam Mata Pelajaran Ke {this.data.jam}</Text>
                <Text style={styles.txttgl}>Terakhir Di Perbarui {this.state.lastupdate ? moment(this.state.lastupdate).format('HH:mm') : "00:00" }</Text>                
            </View>                       
          </View>
          <View style={styles.headerinfo}>   
              <View style={styles.colum}>
                <Text style={{color:'#404040',fontWeight:'bold'}}> Kelas {this.data.Kelas.nama}</Text>
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
    componentDidMount() {
      this._isMounted = true;
      this._isMounted && this._backhand();            
    }

    componentWillUnmount() {  
      this._isMounted = false; 
      BackHandler.removeEventListener('hardwareBackPress', this._goBack);     
    }

    _backhand = () => {
      if(this._isMounted){        
        BackHandler.addEventListener('hardwareBackPress', this._goBack);                 
      }    
    }

    _goBack = () => { 
      //console.warn('NO Pressed')  
      //console.log(this.props.navigation.state.routeName)      
      if(this.props.navigation.isFocused()){
        Alert.alert(
          'Perhatian !',
          'Apakah Waktu Mengajar Anda Selesai ?',        
          [
            {text: 'Tidak', style: 'cancel'},
            {text: 'Ya', onPress: () => this.keluar()},
          ],
          {cancelable: false},
        );
        return true;    
      }     
    };    

    keluar = async () => {        
      await axios.post("http://"+this.state.ipaddress+"/api/absensi/mobile/updatejamkeluar",
      {
          keluar:moment().format('YYYY-MM-DDTHH:mm'),kelas_id:Number(this.data.Kelas.id),jam:Number(this.data.jam)          
      }).then(response => {           
          if(response.data.status == "success")
            { 
              this.props.navigation.goBack();              
            }
         }).catch(error => {           
          
          if(error.message === "Network Error"){            
              this.refs.toast.show('Maaf Jaringan Tidak Tersedia !');
          }
          if(error.response.status == 401){
              this.refs.toast.show(error.response.data.message);               
          }
          
        });
    }

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

    siapapiket = () => {
      Alert.alert(
        'Guru Piket Aktif',
        this.gurupiket.nama,
        [
          {text: 'OK', onPress: () => console.log("cek")},
        ],
        {cancelable: false},
      );      
    }

    _countJumlah = (kode) => {
      var data = this.state.siswa; 
      return data.filter(item => item.kode === kode ).length      
    }

    onRefresh = async () => {       
      //Call the Service to get the latest data
      await this.setState({ siswa:[]},() => this._SiswaData());      
    };    

    _SiswaData = async ()  => {                 
      await axios.get("http://"+this.state.ipaddress+"/api/absensi/getsiswa?kelas="+this.data.Kelas.id+"&tanggal="+moment().format('YYYY-MM-DDTHH:mm')+"&jam="+this.data.jam)
         .then(response => {                                                     
           this.setState({ siswa: response.data.siswa });            
           this.setState({ refreshing: false });
           this.setState({ lastupdate: response.data.tanggal });
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

    _mapelData = (piket) => {
      if (piket === undefined) {var piketid = 0 ; var piketnama = "";}else{piketid = piket.id ; piketnama = piket.nama;}
      var tgl = moment().format('YYYY-MM-DD');  
      axios.post("http://"+this.state.ipaddress+"/api/absensi/store",
      {
          tanggal:moment().format('YYYY-MM-DDTHH:mm'),user_id:Number(this.data.Users.id),kelas_id:Number(this.data.Kelas.id),jam:Number(this.data.jam),kelas:this.data.Kelas.nama,
          mapel:this.data.Users.mapel,mapelguru:this.data.Users.nama,mapel_start:moment(tgl+" "+this.data.jam_start).format('YYYY-MM-DDTHH:mm'),mapel_end:moment(tgl+" "+this.data.jam_end).format('YYYY-MM-DDTHH:mm'),piketid:Number(piketid),piketguru:piketnama

      }).then(response => { 
          this._SiswaData();
          if(response.data.status == "success")
            {                             
              //this.refs.toast.show('Absensi Terupdated Di Server !');              
            }
         }).catch(error => {           
          
          if(error.message === "Network Error"){            
              this.refs.toast.show('Maaf Jaringan Tidak Tersedia !');
          }
          if(error.response.status == 401){
              this.refs.toast.show(error.response.data.message);               
          }
          
        });
    }

    _AbsenData = async ()  => {
      var tgl = moment().format('YYYY-MM-DD');            
      await axios.post("http://"+this.state.ipaddress+"/api/absensi/mobile/store",
      {
          tanggal:moment().format('YYYY-MM-DDTHH:mm'),user_id:Number(this.data.Users.id),kelas_id:Number(this.data.Kelas.id),jam:Number(this.data.jam),kelas:this.data.Kelas.nama,
          mapel:this.data.Users.mapel,mapelguru:this.data.Users.nama,mapel_start:moment(tgl+" "+this.data.jam_start).format('YYYY-MM-DDTHH:mm'),mapel_end:moment(tgl+" "+this.data.jam_end).format('YYYY-MM-DDTHH:mm'),piketguru:this.state.piketguru,
          siswa:JSON.stringify(this.state.siswa)

      }).then(response => { 
          if(response.data.status == "success")
            { 
              this.setState({ lastupdate: response.data.tanggal });
              this.refs.toast.show('Absensi Terperbarui Di Server !');
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
        marginTop:6,        
        fontSize:11,
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