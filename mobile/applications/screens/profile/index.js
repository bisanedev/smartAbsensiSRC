import React from 'react';
import {    
  Image,  
  Text,
  StyleSheet,
  View,
  Alert,
  Dimensions,  
  TouchableOpacity,  
  ScrollView,
  RefreshControl
} from 'react-native';
import AsyncStorage from '@react-native-async-storage/async-storage';
import {widthPercentageToDP as wp, heightPercentageToDP as hp} from 'react-native-responsive-screen';
import Toast, {DURATION} from 'react-native-easy-toast';
import MyStatusBar from '../../components/statusBar';
import Header from './header';
import jwt_decode from 'jwt-decode';
import axios from 'axios';
import Icon from 'react-native-vector-icons/AntDesign';
import { PieChart } from "react-native-chart-kit";
import moment from 'moment';

export default class ProfileScreen extends React.Component {
  
  constructor() {
    super(); 
    this.state = {
      user:'',
      jumlah:'',
      semester:'',
      tahun:'',
      ipserver:'',      
      data:[
        { name: 'Hadir', jumlah: 0, color: 'green', legendFontColor: '#7F7F7F', legendFontSize: 11 },                
        { name: 'Sakit', jumlah: 0, color: '#07ade3', legendFontColor: '#757575', legendFontSize: 11 },
        { name: 'Izin', jumlah: 0, color: 'blue', legendFontColor: '#7F7F7F', legendFontSize: 11 },
        { name: 'Terlambat', jumlah: 0, color: 'orange', legendFontColor: '#7F7F7F', legendFontSize: 11 },
        { name: 'Alpa', jumlah: 0, color: 'red', legendFontColor: '#7F7F7F', legendFontSize: 11 },
      ],
    }

    AsyncStorage.getItem('userToken', (error, token) => {
      axios.defaults.headers.common['Authorization'] = token;
      this.setState({ user: jwt_decode(token) });        
    });

    AsyncStorage.getItem('ipaddress', (error, server) => {
      if (server) {            
        this.setState({
          ipserver: server
        }); 
        this._InsightData();
      }
    });
    //---------- 
  }
    
 
    render() {
      const width = Dimensions.get('window').width;
      const height = 180;
      const chartConfig = {       
        color: (opacity = 1) => `rgba(255, 255, 255, ${opacity})`,
      };

      const graphStyle = {
          marginVertical: -8,          
          ...chartConfig.style
      }

      //------------------  
      return (
        <View style={styles.container}>
          <MyStatusBar backgroundColor="#fff" barStyle="dark-content" />
          <Header/>
          <ScrollView contentContainerStyle={{ width: wp("100%"), alignItems: 'center' }}>
          <View style={this.profileStyle(this.state.user.color)}> 
            <View style={styles.middle}>
              {this.photo(this.state.user)} 
              <Text style={styles.nama}>{this.state.user.nama}</Text>
              <Text style={styles.mapel}>{this.state.user.mapel}</Text> 
            </View>
            <TouchableOpacity style={styles.tombolLogout} onPress={this._ConfirmLogout}>
                <Icon name="logout" size={15}/>
                <Text style={styles.tombolTXT}> KELUAR </Text>
            </TouchableOpacity>
          </View>                    
          
          <View style={styles.konten}> 
          <TouchableOpacity style={styles.tombolRefresh} onPress={this._Refresh}>
              <Icon name="sync" color="#000" size={14}/>
          </TouchableOpacity>
          <Text style={styles.semester}>Tahun Pelajaran {this.state.tahun} Semester {this.state.semester}</Text>         
          <PieChart
            data={this.state.data}
            height={hp("38%")}
            width={wp("100%")}
            chartConfig={chartConfig} 
            backgroundColor="transparent"                   
            accessor="jumlah"
            style={graphStyle}
            paddingLeft="22"
          />         
          </View>
          </ScrollView>
          <Toast ref="toast" defaultCloseDelay={1000} position='top' positionValue={200} opacity={0.8}/>          
        </View>
      );
    }      
    //------------- logout
    _ConfirmLogout = () => {
      Alert.alert(
        'Perhatian !',
        'Yakin Keluar Dari Akun '+this.state.user.nama+' ?',        
        [          
          {text: 'Tidak', style: 'cancel'},
          {text: 'Ya', onPress: () => this._handleLogout()},
        ],
        {cancelable: false},
      );
    }

    _handleLogout = async () => {                  
      delete axios.defaults.headers.common['Authorization']; 
      this.props.logout();
    };

    _Refresh = () => {
      this._InsightData();
    }

    profileStyle = function(options) {
      return {          
        backgroundColor:options,                
        height:hp("47%"),  
        flexDirection: 'row', 
      }
    } 

    photo = (data) => {          
        if(data.foto)
        { 
          return (
            <Image style={styles.pic} source={{uri: "http://"+this.state.ipserver+"/storage/guru/"+data.foto}} />
          )
        }
        else if(data.jenis === "Perempuan")
        { 
          return (
            <Image style={styles.pic} source={{uri: "http://"+this.state.ipserver+"/assets/images/cewek.png"}} />
          )
        }
        else
        {
          return (
            <Image style={styles.pic} source={{uri: "http://"+this.state.ipserver+"/assets/images/cowok.png"}} />
          )
        }
    }

    _InsightData = async ()  => {             
      
      await axios.get("http://"+this.state.ipserver+"/api/info/guru/insight?guru="+this.state.user.id+"&tanggal="+moment().format('YYYY-MM-DDTHH:mm'))
            .then(response => {   
                            
              this.setState({ jumlah: response.data.jumlah });
              this.setState({ semester: response.data.semester });
              this.setState({ tahun: response.data.tahun });
              //modify object json
              let data = [...this.state.data];
              let hadir = data.findIndex(el => el.name === 'Hadir');
              data[hadir] = {...data[hadir], jumlah: response.data.jumlah.hadir};
              let izin = data.findIndex(el => el.name === 'Izin');
              data[izin] = {...data[izin], jumlah: response.data.jumlah.izin};
              let sakit = data.findIndex(el => el.name === 'Sakit');
              data[sakit] = {...data[sakit], jumlah: response.data.jumlah.sakit};
              let terlambat = data.findIndex(el => el.name === 'Terlambat');
              data[terlambat] = {...data[terlambat], jumlah: response.data.jumlah.terlambat};
              let alpa = data.findIndex(el => el.name === 'Alpa');
              data[alpa] = {...data[alpa], jumlah: response.data.jumlah.alpa};
              this.setState({ data });              

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
    
//---------end
}
  
  const styles = StyleSheet.create({
    container: {
      flex:1,
      zIndex:1,
      backgroundColor:'#f2f1f1',    
    },
    tombolRefresh:{
      zIndex:99,           
      borderRadius:20,
      position: 'absolute',
      left: 15,
      top: 30,
      backgroundColor:'#fff',
      padding:5,
      borderColor: '#757575',
      borderWidth: 0.4,
    },
    profil:{
      flex:1,
    },
    semester:{
      fontWeight:'bold',
      fontSize:14,
      marginTop:7,
    },
    nama:{
      marginTop:5,
      fontSize:16,
      color:'white',
      fontWeight:'bold',      
    },
    mapel:{
      fontSize:10,
      color:'white',
      fontWeight:'600',
    },
    middle:{ 
      justifyContent: 'center',      
      alignSelf:'center',
      alignItems: 'center',
      flex:1,
    },
    pic:{
      resizeMode:'contain',
      width:150,
      borderRadius: 5,
      height:150,              
    },
    tombolLogout:{      
      backgroundColor:'#fff',
      borderColor: '#757575',
      borderWidth: 0.4,
      padding:4,
      flexDirection: 'row',      
      marginTop:10,
      position:'absolute',
      bottom:10,
      right:9,
    },
    tombolTXT:{
      fontSize:10,
      fontWeight:'bold',      
    },
    konten:{
      flex:1,
      zIndex:4,                  
      alignSelf:'center',
      alignItems: 'center',          
      backgroundColor:'#fff',
      borderColor: '#e6e6e6',
      borderWidth: 0.7, 
    }
  });  