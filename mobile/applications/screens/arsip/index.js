import React from 'react';
import {      
  Text,   
  ImageBackground,  
  StyleSheet,
  RefreshControl,
  ScrollView,
  View,
} from 'react-native';
import AsyncStorage from '@react-native-async-storage/async-storage';
import {widthPercentageToDP as wp, heightPercentageToDP as hp} from 'react-native-responsive-screen';
import MyStatusBar from '../../components/statusBar';
import {LocaleConfig,Calendar} from 'react-native-calendars';
import axios from 'axios';
import jwt_decode from 'jwt-decode';

LocaleConfig.locales['id'] = {
  monthNames: ['Januari','Februari','Maret','April','Mei','Juni','Juli','Agustus','September','Oktober','November','Desember'],
  monthNamesShort: ['Jan','Feb','Mar','Apr','Mei','Jun','Jul','Agu','Sep','Okt','Nov','Dec'],
  dayNames: ['Minggu','Senin','Selasa','Rabu','Kamis','Jumat','Sabtu'],
  dayNamesShort: ['Min','Sen','Sel','Rab','Kam','Jum','Sab'],
  today: 'Hari ini'
};

LocaleConfig.defaultLocale = 'id';


export default class ArsipScreen extends React.Component {

  constructor(props) {
    super(props);
    this.state = {
      tanggal:[],
      users:[],      
      ipserver:'',
      refreshing: true,
      marked:null,
    };

    AsyncStorage.getItem('userToken', (error, token) => {
      axios.defaults.headers.common['Authorization'] = token;
      this.setState({ user: jwt_decode(token) });       
    });

    AsyncStorage.getItem('ipaddress', (error, server) => {
      if (server) {            
        this.setState({
          ipserver: server
        });                  
        this._GetTanggal();         
      }
    }); 

  }
  
    render() {         
      return (
        <View style={styles.container}>
          <MyStatusBar backgroundColor="#fff" barStyle="dark-content" />                    
          <ImageBackground source={require('../../assets/pattern.jpg')} style={styles.background}>
          <View style={styles.header}>
            <Text style={styles.titleText}>Arsip Absensi</Text>                  
          </View>          
          </ImageBackground>
          <ScrollView refreshControl={
            <RefreshControl refreshing={this.state.refreshing} onRefresh={this.onRefresh.bind(this)} />
          }> 
          <View style={styles.calendarkonten}>
          <Calendar
            onDayPress={ ( day ) => this.onDayPress( day ) }            
            style={styles.calendar}            
            firstDay={1}            
            markedDates={this.state.marked}
            theme={{
              todayTextColor: '#3D7D7D',
              arrowColor: '#757575',
            }}                     
          /> 
          </View>                          
          </ScrollView>              
        </View>
      );
    }    

    onDayPress(day) {
      const data = this.state.tanggal

      if(data.some(item => item.Tgl === day.dateString)){
        //alert(day.dateString);    
        this.props.navigation.navigate('Detail', {id:this.state.user.id,tanggal:day.dateString})    
      }          
      
    }

    onRefresh = async () => { 
      await this.setState({ tanggal:[],marked:[]});      
      //Call the Service to get the latest data
      this._GetTanggal();
    };   
    
    _GetTanggal = async () => {         
    await axios.get("http://"+this.state.ipserver+"/api/absensi/listtanggal?guru="+this.state.user.id)
      .then(response => {   

       const data = response.data.data;
       this.setState({ tanggal: data });
       this.setState({ refreshing: false });  
       // create marked  
       var obj = data.reduce((c, v) => Object.assign(c, {[v.Tgl]: {selected: true,selectedColor: '#757575'}}), {});
       this.setState({ marked : obj});

      }).catch(error => {           
       
       if(error.message === "Network Error"){            
           this.refs.toast.show('Maaf Jaringan Tidak Tersedia !');
       }
       if(error.response.status == 401){
           this.refs.toast.show(error.response.data.message);               
       }
       
      });          
    };
  }
  
  const styles = StyleSheet.create({
    container: {
      flex:1,
      backgroundColor:'white',
    },
    calendar: {                                       
      marginLeft:10,
      marginRight:10,
      width:wp("100%"),
    },
    header:{      
      height:55,           
      width:'100%',
      borderColor: '#e6e6e6',
      borderWidth: 0.7,
      flexDirection:'row',          
      backgroundColor:'white',
      opacity: 0.9,
      elevation:3,        
      alignItems: 'center',
      justifyContent: 'center', 
    },    
    calendarkonten:{
      textAlign:'center',      
      alignItems: 'center',
      justifyContent: 'center', 
      backgroundColor:'white',      
    },
    titleText:{
      marginRight:12,
      fontSize:19,
      fontWeight:'bold',
      textTransform: 'uppercase',
    },
    background: {           
      resizeMode: 'stretch',
      height:55,      
      width:'100%',        
      flex:1,                              
    }
  });