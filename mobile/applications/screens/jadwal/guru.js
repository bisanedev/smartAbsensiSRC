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

export default class GuruScreen extends React.Component {

    constructor(props) {
      super(props);          
      this.state = {
        guru:[],        
        refreshing: true,
        ipserver:'',
        piketguru:'',        
      };     
      AsyncStorage.getItem('ipaddress', (error, server) => {
        if (server) {            
          this.setState({
            ipserver: server
          });
          this._GuruData();
          this._getPiket();
          this.RenderWebsocket();
        }
      });            
      this.kode = this.props.route.params.id;    
    }

  
    render() {  
      
      return (
        <View style={styles.container}>
          <MyStatusBar backgroundColor="#fff" barStyle="dark-content" />
          <Header kembali={this._goBack}/>
          <ScrollView  refreshControl={
            <RefreshControl refreshing={this.state.refreshing} onRefresh={this.onRefresh.bind(this)} />
          }>
          
          { this.state.guru.sort((a, b) => this.dayOfWeekAsInt(a.hari) > this.dayOfWeekAsInt(b.hari)).map((item,i) => (
              <Body data={item} key={i} gurupiket={this.state.piketguru}/>
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

    _goBack = () => {            
        this.props.navigation.goBack();
    };

    _getPiket = async () => {
      await AsyncStorage.getItem('userToken', (error, token) => {        
        this.setState({ user: jwt_decode(token) });        
      }); 
      axios.get("http://"+this.state.ipserver+"/api/users/"+this.state.user.id)
      .then(response => { 
        if(response.data.status == "success")
        {
          const data = response.data.data;
          this.setState({ piketguru: data })        
        }                 
      }).catch(error => { 
        if(error.message === "Network Error"){            
          this.refs.toast.show('Maaf Jaringan Tidak Tersedia !');
        }
      });      
    };

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
      
      axios.get("http://"+this.state.ipserver+"/api/jadwal?guru="+this.kode )
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