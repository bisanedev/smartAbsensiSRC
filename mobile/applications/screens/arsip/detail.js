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

export default class ArsipDetailScreen extends React.Component {

    constructor(props) {
      super(props);          
      this.state = {
        absen:[],        
        refreshing: true,
        ipserver:'',
      };   
            
      this.kode = this.props.route.params.id;
      this.tanggal = this.props.route.params.tanggal;      
    }

    componentDidMount(){
        AsyncStorage.getItem('ipaddress', (error, server) => {
            if (server) {            
              this.setState({
                ipserver: server
              });
              this._AbsenData();
            }
        });
    }

    render() {  
      
      return (
        <View style={styles.container}>
          <MyStatusBar backgroundColor="#fff" barStyle="dark-content" />
          <Header kembali={this._goBack} tanggal={this.tanggal}/>
          <ScrollView  refreshControl={
            <RefreshControl refreshing={this.state.refreshing} onRefresh={this.onRefresh.bind(this)} />
          }>
          
          { this.state.absen.sort((a, b) =>  a.jam > b.jam).map((item,i) => (
              <Body data={item} key={i}/>
          ))}
            
          </ScrollView>          
          <Toast ref="toast" defaultCloseDelay={1000} position='top' positionValue={200} opacity={0.8}/>
        </View>
      );
    }   
  
    onRefresh() {
      this.setState({ absen: [] },() => this._AbsenData());      
    }

    _goBack = () => {            
        this.props.navigation.goBack();
    };

    _AbsenData = ()  => {             
      
      axios.get("http://"+this.state.ipserver+"/api/absensi/mobile?guru="+this.kode+"&tanggal="+this.tanggal )
         .then(response => {   

          const data = response.data.data;
          this.setState({ absen: data,refreshing: false });                    
          
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
});