import React from 'react';
import { StyleSheet,View,Text,TouchableOpacity } from 'react-native';
import { withNavigation } from '@react-navigation/compat';
import AsyncStorage from '@react-native-async-storage/async-storage';
import moment from 'moment';

class Body extends React.Component {
    constructor() {
        super();     
        this.state = {ipaddress: '',keluar:false};       
    }

    componentDidMount(){
        AsyncStorage.getItem('ipaddress', (error, result) => {
            if (result) {            
                this.setState({
                    ipaddress: result
                });
            }
        });
    }

    render() {
        let data = this.props.data;
        let mapelnote;
        
        if (data.mapel_note){
            mapelnote = <Text style={styles.txt}>Guru Piket {data.mapel_note}</Text>
        }
        
        if (data.keluar == "0001-01-01T00:00:00Z"){
            this.state.keluar = false;
        }else{
            this.state.keluar = true;
        }
        //-----------
        return (
            <View style={styles.container}>
                <View style={styles.card}>                
                <View style={styles.konten}>                                                            
                    <Text style={styles.txt}>Mata Pelajaran {data.mapel} [ {data.kelas} ]</Text>                    
                    <Text style={styles.txt}>Jam {moment(data.mapel_start).format('HH:mm')} ~ {moment(data.mapel_end).format('HH:mm')} [{moment.duration(moment(data.mapel_end).diff(moment(data.mapel_start))).asMinutes()} Menit] , Mata Pelajaran Ke {data.jam}</Text>                    
                    {mapelnote}                                          
                    <Text style={styles.kelas}>Masuk Kelas {moment(data.tanggal).format('HH:mm')} ~ Selesai {this.state.keluar ? moment(data.keluar).format('HH:mm') : "00:00"} [{this.state.keluar ? moment.duration(moment(data.keluar).diff(moment(data.tanggal))).asMinutes() : "0"} Menit]</Text>
                    <Text style={styles.edit}>Terakhir Di Edit {moment(data.updated_at).format('DD/MM/YYYY HH:mm:ss')}</Text>
                    <Text style={styles.kehadiran}>( Hadir : {this._countJumlah('h')} , Sakit: {this._countJumlah('s')} , Izin: {this._countJumlah('i')} , Terlambat: {this._countJumlah('t')} , Alpa: {this._countJumlah('a')} )</Text>
                </View>
                <View style={styles.lastEdit}>
                    
                </View>
                <TouchableOpacity style={styles.tombol} onPress={() => { this.props.navigation.navigate('Edit', {dataEdit: data});}}>
                    <Text style={styles.tombolmasuk}> EDIT </Text>
                </TouchableOpacity>                                                                
                </View>                              
            </View>
        );
    }           
//-----------
_countJumlah = (kode) => {
    //var data = JSON.parse(this.props.data.Siswa); 
    var data =this.props.data.Siswa; 
    return data.filter(item => item.kode === kode ).length      
}
//------------
}

const styles = StyleSheet.create({
    container: {
      flex: 1,          
    },  
    tomboldis:{
        display: 'none',       
    },
    tombol:{
        padding:4,
        position:'absolute',
        bottom:-1,
        right:-1,
        borderRadius:0,        
        backgroundColor:'#e6e6e6',                  
    },
    tombolmasuk:{
        fontWeight:'bold',
        fontSize:11,
    },
    card:{
        backgroundColor:"#757575",
        flexDirection:'row', 
        justifyContent:'space-between', 
        alignItems:'center',                 
        marginTop:15,
        marginBottom:5,          
        marginLeft:10,
        marginRight:10,
        borderColor: '#757575',
        borderWidth: 0.4,  
    },  
    konten:{                     
        flex: 1,
        marginTop:25,
        marginBottom:25,
        alignItems: 'center',        
    },  
    kelas:{
        marginTop:8,        
        fontWeight:'bold',
        fontSize:9,
        color:'#f3f3f3',
    },
    edit:{
        marginTop:4,
        fontStyle:'italic',
        fontWeight:'bold',
        fontSize:9,
        color:'#f3f3f3',
    },
    txt:{
        textAlign:'center', 
        marginTop:3,        
        fontSize:12,
        color:'#fefefe',
    },
    kehadiran:{
        fontSize:10,
        textAlign:'center', 
        marginTop:11,
        fontWeight:'bold',
        color:'#fefefe',
    }

});

export default withNavigation(Body);