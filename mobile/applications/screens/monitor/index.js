import React from 'react';
import {      
  Text,
  StyleSheet,
  View,  
  ScrollView,
  RefreshControl,
  TouchableOpacity,  
} from 'react-native';
import AsyncStorage from '@react-native-async-storage/async-storage';
import MyStatusBar from '../../components/statusBar';
import Header from './header';
import axios from 'axios';
import jwt_decode from 'jwt-decode';
import Icon from 'react-native-vector-icons/MaterialIcons';

export default class MonitorScreen extends React.Component {
    constructor(props) {
        super(props); 
        this.state = {
            jumlahsiswa:"",
            refreshing: true,
            absen:[],
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
          this.RenderWebsocket();
          this.AbsenData();
        }
      });
    //----end constructor
    }

    render() {
      //------------------  
      return (
        <View style={styles.container}>
            <MyStatusBar backgroundColor="#fff" barStyle="dark-content" />
            <Header/>            
            <ScrollView refreshControl={
            <RefreshControl refreshing={this.state.refreshing} onRefresh={this.onRefresh.bind(this)} />
            }>
            <View style={this.cardStyle("#757575")}>
                <Text style={styles.number}>{this.state.jumlahsiswa-(this.countJumlah('h')+this.countJumlah('t')+this.countJumlah('i')+this.countJumlah('s')+this.countJumlah('a'))}</Text>
                <Text style={styles.title}>Siswa Tidak Absensi</Text>
            </View>
            <View style={this.cardStyle("#008000")}>                
                <Text style={styles.number}>{this.countJumlah('h')}</Text>
                <Text style={styles.title}>Siswa Hadir</Text>
            </View>
            <View style={this.cardStyle("#07ade3")}>
                <Text style={styles.number}>{this.countJumlah('s')}</Text>
                <TouchableOpacity style={styles.more} onPress={() => {this.ListDetail('s','SISWA SAKIT')}}>
                    <Text style={styles.title}>Siswa Sakit</Text>
                    <Icon name="more-vert" size={22} color="#fff" style={styles.icon}/>
                </TouchableOpacity>
            </View>
            <View style={this.cardStyle("#0716e3")}>
                <Text style={styles.number}>{this.countJumlah('i')}</Text>
                <TouchableOpacity style={styles.more} onPress={() => {this.ListDetail('i','SISWA IZIN')}}>
                    <Text style={styles.title}>Siswa Izin</Text>
                    <Icon name="more-vert" size={22} color="#fff" style={styles.icon}/>
                </TouchableOpacity>                
            </View>
            <View style={this.cardStyle("#fc9403")}>
                <Text style={styles.number}>{this.countJumlah('t')}</Text>
                <TouchableOpacity style={styles.more} onPress={() => {this.ListDetail('t','SISWA TERLAMBAT')}}>
                    <Text style={styles.title}>Siswa Terlambat</Text>
                    <Icon name="more-vert" size={22} color="#fff" style={styles.icon}/>
                </TouchableOpacity>                
            </View>
            <View style={this.cardStyle("#e30707")}>
                <Text style={styles.number}>{this.countJumlah('a')}</Text>
                <TouchableOpacity style={styles.more} onPress={() => {this.ListDetail('a','SISWA ALPA')}}>
                    <Text style={styles.title}>Siswa Alpa</Text>
                    <Icon name="more-vert" size={22} color="#fff" style={styles.icon}/>
                </TouchableOpacity>                
            </View>
            </ScrollView>            
        </View>
      );
      //------------------
    }
    //function or method
    RenderWebsocket = () => {        
        var ws = new WebSocket("ws://"+this.state.ipserver+"/websocket/absen/ws");
        ws.onopen = () => {
            // connection opened
            console.log("trigger store Konek");
            ws.onmessage = ({data}) => {          
                this.AbsenData();
            };
        };
        var ws2 = new WebSocket("ws://"+this.state.ipserver+"/websocket/absenupdate/ws");
        ws2.onopen = () => {
            // connection opened
            console.log("trigger updated Konek");
            ws2.onmessage = ({data}) => {          
                this.AbsenData();
            };
        };
    }

    countJumlah(kode){
        var data = this.state.absen;
        return data.filter(item => item.kode === kode ).length      
    }

    ListDetail = (kode,title) => {
        var data = this.state.absen.filter(item => item.kode === kode );         
        this.props.navigation.navigate('Detail', {monitor:data,title:title});        
    }

    onRefresh() {
        //Call the Service to get the latest data
        this.setState({ absen: [] }, () => this.AbsenData());                      
    }

    AbsenData = ()  => {
        axios.get("http://"+this.state.ipserver+"/api/info/absen/harian")
        .then(response => {
            this.setState({ jumlahsiswa: response.data.jumlah });
            this.setState({ absen: response.data.absen });
            this.setState({ refreshing: false });            
        }).catch(error => {this.state.absen = [];this.state.jumlahsiswa = "";})
    }

    cardStyle = function(options) {
        return {          
          backgroundColor:options,
          justifyContent:'center', 
          alignItems:'center',                 
          marginTop:11,
          height:120,
          marginBottom:11,          
          marginLeft:10,
          marginRight:10,
          borderColor: '#757575',
          borderWidth: 0.4,                          
        }
    }      
//---------end
}

const styles = StyleSheet.create({
    container: {
      flex:1,
      zIndex:1,
      backgroundColor:'#f2f1f1',      
    },
    title:{
        fontWeight:'normal',
        fontSize:20,
        color:"#fff",
        marginBottom:8,
        textTransform: 'uppercase',
    },
    icon:{
        marginTop:3,
    },
    more:{
        flexDirection:'row'
    },
    number:{
        fontWeight:'bold',
        fontSize:54,        
        color:"#fff",       
    },
});