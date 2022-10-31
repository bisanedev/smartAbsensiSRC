import React from 'react';
import { View , StyleSheet,Text ,Image ,TextInput,TouchableOpacity,ImageBackground,Alert } from 'react-native';
import MyStatusBar from '../../components/statusBar';
import axios from 'axios';
import AsyncStorage from '@react-native-async-storage/async-storage';
import Icon from 'react-native-vector-icons/Ionicons';
import {widthPercentageToDP as wp, heightPercentageToDP as hp} from 'react-native-responsive-screen';
import Toast, {DURATION} from 'react-native-easy-toast';

export default class SignInScreen extends React.Component {
    constructor(props) {
        super(props);
        this.params = this.props.route.params;    
        this.state = { 
            username:"",        
            password:"",
            inputError:"",
            disableSubmit:false,
            hidePassword:true,
            setIcon:"eye-outline",
            ipserver:"",
            ipaddress:false
        }
    }

    componentDidMount() { 
        AsyncStorage.getItem('ipaddress', (error, result) => {
            if (result) {            
                this.setState({ipserver: result,ipaddress:true});
            }
        });
    }
    
    componentWillUnmount() {
    
    }

    render() {
        const { username,password,disableSubmit,ipserver,ipaddress,inputError } = this.state;
        if(ipaddress === false){
            //halaman simpan ip server
            return (
                <View style={styles.container}>
                    <MyStatusBar backgroundColor={global.BarColor} barStyle="dark-content" />
                    <ImageBackground source={require('../../assets/background_dot.png')} resizeMode="repeat" style={styles.background}/>             
                    <TextInput
                        placeholder="ex 192.168.1.2:9000"
                        autoCapitalize="none"                        
                        returnKeyType="done"
                        value={ipserver}
                        onChangeText={(ipserver) => this.setState({ ipserver })}                        
                        style={styles.input}
                    />
                    <TouchableOpacity style={styles.tombol} onPress={this.simpanIPServer.bind(this)}>
                        <Text style={styles.tombolTxt}>Koneksi Ke SERVER</Text>
                    </TouchableOpacity>                    
                    <Text style={styles.error}>{this.state.err}</Text>
                </View>
            );
            //end halaman simpan ip server
        }else{
            //halaman login
            return (
                <View style={styles.container}>
                    <MyStatusBar backgroundColor={global.BarColor} barStyle="dark-content" />   
                    <ImageBackground source={require('../../assets/login.png')} style={{...styles.background,opacity: 0.2}}/>              
                    <Image source={{uri: 'http://'+ipserver+'/storage/logo.png'}} style={styles.logo}/>
                     <TextInput
                       placeholder="Username"
                       autoCapitalize="none"
                       autoCompleteType="username"
                       returnKeyType="next"
                       value={username}
                       onChangeText={(username) => this.setState({username: username.split(" ").join("") })}
                       style={styles.input}
                     />
                    {inputError.username && (<Text style={styles.error}>{inputError.username}</Text>)}
                    <View style={styles.inputPasswordLayout}>
                    <TextInput
                        placeholder="Password"
                        autoCapitalize="none"
                        autoCompleteType="password"
                        returnKeyType="done"
                        value={password}
                        onChangeText={(password) => this.setState({ password })}
                        secureTextEntry={this.state.hidePassword}
                        style={styles.inputPassword}
                    />
                    <Icon name={this.state.setIcon} style={styles.icon} size={20} onPress={() => this.changeIcon()}/>
                    </View>
                    {inputError.password && (<Text style={styles.error}>{inputError.password}</Text>)}
                    <TouchableOpacity style={styles.tombol} onPress={this.login.bind(this) } disabled={disableSubmit}>
                        <Text style={{...styles.tombolTxt,color:disableSubmit? "lightgrey" : "white"}}>Authentikasi</Text>
                    </TouchableOpacity>
                        {inputError.network && (<Text style={styles.error}>{inputError.network}</Text>)}
                        {inputError.notFound && (<Text style={styles.error}>{inputError.notFound}</Text>)}
                    <TouchableOpacity style={styles.tombolganti} onPress={() => this.confirmGanti()}>
                        <Text style={styles.tombolexit}> GANTI SERVER </Text>
                    </TouchableOpacity>
                    <Toast ref="toast" defaultCloseDelay={1000} position='top' positionValue={200} opacity={0.8}/> 
                </View>
            );
            //end halaman login
        }        
      }
// --------------- function
  changeIcon = () => {
    const { setIcon } = this.state;
    setIcon === "eye-outline"
      ? this.setState({ setIcon:"eye-off-outline",hidePassword:false })
      : this.setState({ setIcon:"eye-outline",hidePassword:true })
  }

  simpanIPServer = () => {    
    const { ipserver } = this.state;
    if(ipserver === ""){
        this.refs.toast.show("Input Ip SERVER kosong");  
        return false;
    }
    axios.get("http://"+ipserver+"/api/checkauth", null, null).then(response => { 
        if(response.data.status == "success")
        {
           //kosongi
        }  
    }).catch(error => {   
        if(error.response.status == 401){
            //ke halaman login
            this.setState({ ipaddress:true },() => {
                AsyncStorage.setItem('ipaddress',ipserver);
            });    
        }                  
        if(error.message === "Network Error"){            
            this.setState({err: "Maaf Jaringan "+ipserver+" Tidak Tersedia !" });
        }                 
    });       
  };

  confirmGanti = () => {
    Alert.alert(
      'Perhatian !',
      'Yakin Ganti Konfigurasi IP Address Server?',        
      [          
        {text: 'Tidak', style: 'cancel'},
        {text: 'Ya', onPress: () => 
            this.setState({ipaddress:false},() => {AsyncStorage.clear()})
        },
      ],
      {cancelable: false},
    );
  }

  login = () => {   
    const { username,password,ipserver } = this.state;  
    this.setState({ disableSubmit:true });  

    axios.post('http://'+ipserver+'/login',{
        username: username , password: password
    }).then(response => {
      if(response.data.status == "success")
      {           
       var userToken = response.data.token;
       var sekolah = response.data.sekolah;
       axios.defaults.headers.common['Authorization'] = userToken;               
       const items = [['userToken', userToken], ['sekolah', sekolah]]
       AsyncStorage.multiSet(items, () => {
        this.props.login({ userToken });
       });     
      }
    }).catch(error => {    
      this.setState({ disableSubmit:false });
      if(error.message === "Network Error"){      
        this.setState({ inputError:{network:"Network Error"} });
      }
      if(error.response.status == 401){
        this.setState({ inputError:{notFound:error.response.data.message} });      
      }
      if(error.response.status == 400){
        this.setState({ inputError:error.response.data.message });
      }
    }); 
  }
  //end
}


const styles = StyleSheet.create({
    container: {       
      alignItems:"center",    
      justifyContent:"center",
      flex:1
    },
    background: {   
        resizeMode: 'cover',  
        position: 'absolute',
        top: 0,
        left: 0,
        bottom: 0,
        right: 0      
      },
    error:{ 
        marginTop:10,     
        color:'red',
        fontSize:12,
        fontStyle: 'italic',
        fontWeight: 'bold',
        alignSelf: 'center',    
      },
    input:{  
      width: wp('80%'),
      height: hp('6.7%'),
      marginTop:8,
      borderColor:'lightgrey',
      borderStyle:'solid',
      borderWidth: 1,
      padding:10,
      borderRadius: 3,
      backgroundColor: "#fff", 
    },
    inputPasswordLayout:{  
      width: wp('80%'),
      height: hp('6.7%'),   
      marginTop:8,
      borderColor:'lightgrey',
      borderStyle:'solid',
      borderWidth: 1,    
      borderRadius: 3,
      backgroundColor: "#fff", 
    },
    inputPassword:{
      padding:hp("1.8%"),
    },
    tombol:{    
      marginTop:15,
      width:wp("80%"),
      height:hp("6.8%"),    
      backgroundColor:"#757575",    
      alignItems:"center",
      justifyContent:"center",
      borderRadius:3,    
    },
    tombolganti:{
        padding:4,
        position:'absolute',
        bottom:-1,
        right:-1,
        borderRadius:0,        
        backgroundColor:'#fff',
        borderColor:'#e6e6e6',
        borderWidth:0.7,
    },
    tombolTxt:{
      color:"#fff",
      textAlign:"center",
      fontWeight:"600",
      fontSize:13,
    },
    tombolexit:{
        fontWeight:'bold',
        fontSize:11,
    },
    logo: {
      width: wp('60%'),
      height: hp('35%'),
      resizeMode: 'stretch',
      marginBottom:15,
    },
    error:{
      marginTop:8,
      fontSize:12,
      fontWeight:'bold',
      color:'red',
    },
    icon:{
      color:'lightgrey',
      position:'absolute',
      right:10,    
      top:'27%',
    }
  });