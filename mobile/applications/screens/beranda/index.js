import React from 'react';
import {     
  StyleSheet,
  View,  
  ScrollView,
  RefreshControl, 
} from 'react-native';
import AsyncStorage from '@react-native-async-storage/async-storage';
import Toast, {DURATION} from 'react-native-easy-toast';
import MyStatusBar from '../../components/statusBar';
import Header from './header';
import Body from './body';
import axios from 'axios';
import jwt_decode from 'jwt-decode';

export default class BerandaScreen extends React.Component {

    constructor() {
      super();          
      this.state = {
        guru:[],        
        refreshing: true,
        ipserver:'',
        sekolah:'',
      };
      AsyncStorage.getItem('userToken', (error, token) => {
        axios.defaults.headers.common['Authorization'] = token;
        this.setState({ user: jwt_decode(token) });        
      });
      AsyncStorage.getItem('sekolah', (error, data) => {        
        this.setState({ sekolah: data });        
      });
      AsyncStorage.getItem('ipaddress', (error, server) => {
        if (server) {            
          this.setState({
            ipserver: server
          });
          this._GuruData();
          this.RenderWebsocket();
        }
      });      
       
    }   
  
    render() {  
      
      return (
        <View style={styles.container}>
          <MyStatusBar backgroundColor="#fff" barStyle="dark-content" />
          <Header logo={"http://"+this.state.ipserver+"/storage/logo.png"} nama={this.state.sekolah} />
          <ScrollView  refreshControl={
            <RefreshControl refreshing={this.state.refreshing} onRefresh={this.onRefresh.bind(this)} />
          }>
          
          { this.state.guru.sort((a, b) => this.dayOfWeekAsInt(a.hari) > this.dayOfWeekAsInt(b.hari)).map((item,i) => (
              <Body data={item} key={i}/>
          ))}
            
          </ScrollView>          
          <Toast ref="toast" defaultCloseDelay={1000} position='top' positionValue={200} opacity={0.8}/>
        </View>
      );
    }   

    dayOfWeekAsInt = function(dayIndex) {      
      if(dayIndex == "Senin"){return 0;}
      if(dayIndex == "Selasa"){return 1;}
      if(dayIndex == "Rabu"){return 2;}
      if(dayIndex == "Kamis"){return 3;}
      if(dayIndex == "Jumat"){return 4;}
      if(dayIndex == "Sabtu"){return 5;}      
      if(dayIndex == "Minggu"){return 6;}
    };

    onRefresh() {
      this.setState({ guru: [] });      
      //Call the Service to get the latest data
      this._GuruData();
    }

    //function or method
    RenderWebsocket = () => {        
      var ws2 = new WebSocket("ws://"+this.state.ipserver+"/websocket/absenupdate/ws");
      ws2.onopen = () => {
          // connection opened          
          ws2.onmessage = ({data}) => {                    
            this._GuruData();
          };
      };
    }

    _GuruData = async ()  => {             
      
      axios.get("http://"+this.state.ipserver+"/api/jadwal?guru="+this.state.user.id)
         .then(response => {   

          const data = response.data.data;
          this.setState({ guru: data });
          this.setState({ refreshing: false });
          //console.log(JSON.stringify(data));
          
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
    backgroundColor:'#f2f1f1',       
  },
});