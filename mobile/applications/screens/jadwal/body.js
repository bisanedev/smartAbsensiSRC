import React from 'react';
import { StyleSheet,View,Image,Text,TouchableOpacity } from 'react-native';
import { withNavigation } from '@react-navigation/compat';
import AsyncStorage from '@react-native-async-storage/async-storage';
import moment from 'moment';
import Dialog from "react-native-dialog";
import axios from 'axios';
import Toast, {DURATION} from 'react-native-easy-toast';
import "moment/locale/id";

class Body extends React.Component {
    constructor(props) {
        super(props);     
        this.state = {ipaddress: '',dialogVisible:false,password:''};
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
        const cewek = <Image style={styles.pic} source={{uri: "http://"+this.state.ipaddress+"/assets/images/cewek.png"}} />;
        const cowok = <Image style={styles.pic} source={{uri: "http://"+this.state.ipaddress+"/assets/images/cowok.png"}} />;
        const photo = <Image style={styles.pic} source={{uri: "http://"+this.state.ipaddress+"/storage/guru/"+data.Users.foto}} />;
        let profile;
        if(data.Users.foto){ profile = photo }
        else if(data.Users.jenis === "Perempuan"){profile = cewek}else{profile = cowok}

        return (
            <View style={styles.container}>
                <View style={this.cardStyle(data.Users.color,data.hari)}>
                {profile}
                <View style={styles.konten}>
                    <Text style={styles.txtnama}>{data.Users.nama}</Text>                                        
                    <Text style={styles.txt}>Mata Pelajaran {data.Users.mapel}</Text>
                    <Text style={styles.txt}>Hari {data.hari} , Jam {data.jam_start} ~ {data.jam_end}</Text>
                    <Text style={styles.txtjam}>Jam Mata Pelajaran Ke {data.jam}</Text>                    
                </View>
                <TouchableOpacity style={ this.cekhari(data.hari,data.jam_start,data.jam_end) ? styles.tomboldis : styles.tombol} onPress={this.TampilkanDialog}>
                    <Text style={styles.tombolmasuk}> BUKA </Text>
                </TouchableOpacity>                                
                <Text style={styles.txtkelas}>{data.Kelas.nama}</Text>                  
                </View>
                <Dialog.Container visible={this.state.dialogVisible}>
                    <Dialog.Title>Konfirmasi Password Piket</Dialog.Title>                    
                    <Dialog.Input style={styles.inputBox} autoCapitalize="none" onChangeText={(password) => this.setState({password})} />                              
                    <Dialog.Button label="Batal" onPress={this.handleCancel} />
                    <Dialog.Button label="Autentikasi" onPress={this.handlePassword} />
                </Dialog.Container>
                <Toast ref="toast" defaultCloseDelay={1000} position='top' positionValue={200} opacity={0.8}/>
            </View>
        );
    }  

    TampilkanDialog = () => {
        this.setState({ dialogVisible: true });
    };

    handlePassword = () => {
    let data = this.props.data;   
    // ...Your logic
    axios.post("http://"+this.state.ipaddress+"/api/sistem/piket/auth",
      {
        id:1,password_piket:this.state.password          
      }).then(response => {           
          if(response.data.status == "success")
            { 
                this.setState({ dialogVisible: false });
                this.props.navigation.navigate('AbsenJadwal', {data:data,disablepiket:false,piket:this.props.gurupiket})            
            }
         }).catch(error => {           
          this.setState({ dialogVisible: false });
          if(error.message === "Network Error"){            
              this.refs.toast.show('Maaf Jaringan Tidak Tersedia !');
          }
          if(error.response.status == 401){
              this.refs.toast.show(error.response.data.message);               
          }
          
    });
    
    };

    handleCancel = () => {
        this.setState({ dialogVisible: false });
    };

    cekhari = function(hari,start,end) {
        let tgl = moment().format('YYYY-MM-DD');
        let mulai = moment(tgl+" "+start).format('YYYY-MM-DD HH:mm');
        let akhir = moment(tgl+" "+end).format('YYYY-MM-DD HH:mm');
        let now = moment().format('d');    
        
        if(hari == this.dayOfWeekAsString(now)  && moment().isBetween(mulai,akhir) ){
            return false;
        }else{
            return true;
        } 
    }        

    dayOfWeekAsString = function(dayIndex) {
        return ["Minggu","Senin","Selasa","Rabu","Kamis","Jumat","Sabtu"][dayIndex];
    }

    cardStyle = function(options,hari) {
        if(moment().format('dddd') == hari){
            return {          
                backgroundColor:options,
                flexDirection:'row', 
                justifyContent:'space-between', 
                alignItems:'center',                 
                marginTop:10,
                marginBottom:10,          
                marginLeft:10,
                marginRight:10,
                borderColor: '#757575',
                borderWidth: 0.4,                                        
              }
        }else{
            return {          
                backgroundColor:options,
                flexDirection:'row', 
                justifyContent:'space-between', 
                alignItems:'center',                 
                marginTop:10,
                marginBottom:10,          
                marginLeft:10,
                marginRight:10,
                borderColor: '#757575',
                borderWidth: 0.4,
                opacity: .7                        
            }
        }
    }    

}

const styles = StyleSheet.create({
    container: {
      flex: 1,          
    },
    inputBox:{
        borderColor: '#757575',
        borderWidth: 0.4,  
        backgroundColor: '#F7F7F8',
    },
    pic:{        
        width:80,
        height:120,        
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
    txtkelas:{
        position:'absolute',
        fontWeight:'bold',
        padding:3,
        color:'#fff',
        top:-4,
        right:0,
        fontSize:15,
    },
    konten:{        
        marginLeft:-20,
        alignItems: 'center',        
        flex: 1,
    },
    txt:{
        marginTop:3,        
        fontSize:11,
        color:'#fefefe',
    },
    txtjam:{
        marginTop:15,        
        fontSize:12,
        color:'#fefefe',
    },
    txtnama:{        
        fontWeight:'bold',
        fontSize:18,
        marginBottom:10,        
        color:'#fefefe',
    }

});

export default withNavigation(Body);