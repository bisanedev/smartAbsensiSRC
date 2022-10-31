import React from 'react';
import {     
  StyleSheet,
  View,
  Image,
  Text,  
  TouchableOpacity,  
  ScrollView,
  ImageBackground,  
  RefreshControl, 
} from 'react-native';
import AsyncStorage from '@react-native-async-storage/async-storage';
import Toast, {DURATION} from 'react-native-easy-toast';
import MyStatusBar from '../../components/statusBar';
import axios from 'axios';
import SearchInput, { createFilter } from 'react-native-search-filter';
import jwt_decode from 'jwt-decode';
import Icon from 'react-native-vector-icons/Ionicons';

const KEYS_TO_FILTERS = ['nama', 'mapel'];

export default class JadwalScreen extends React.Component {

    constructor(props) {
      super(props);      
      this.state = {
        users:[],        
        refreshing: true,
        ipserver:'',
        searchTerm: '',       
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
          this.fetchGuru(); 
        }
      });            
      
    }
  
    render() {
      const data = this.state.users;
      const filteredEmails = data.filter(createFilter(this.state.searchTerm, KEYS_TO_FILTERS))      
      //------
      return (
        <View style={styles.container}>
          <MyStatusBar backgroundColor="#fff" barStyle="dark-content" />            
          <ImageBackground source={require('../../assets/pattern.jpg')} style={styles.background}>
            <View style={styles.header}>
            <Icon name="ios-search" style={styles.search} size={30}/>
            <View style={styles.caribox}>
            <SearchInput 
            onChangeText={(term) => { this.searchUpdated(term) }}                     
            style={styles.searchInput}
            placeholder="Ketik Nama / Mapel"       
            />
            </View>
            </View>
          </ImageBackground>
                    
          <ScrollView refreshControl={
            <RefreshControl refreshing={this.state.refreshing} onRefresh={this.onRefresh.bind(this)} />
          }>
            {filteredEmails.map(data => {
              return (
                <TouchableOpacity onPress={()=> { this.props.navigation.navigate('Guru', {id:data.id});}} key={data.id} style={styles.DataItem}>
                  {this.renderPhoto(data)}
                  <View style={styles.konten}>
                    <Text>{data.nama}</Text>
                    <Text style={styles.mapel}>{data.mapel}</Text>
                  </View>
                </TouchableOpacity>
              )
            })}
          </ScrollView>                   
          <Toast ref="toast" defaultCloseDelay={1000} position='top' positionValue={200} opacity={0.8}/>
        </View>
      );
    }   

    renderPhoto(data) {      
      const cewek = <Image style={styles.pic} source={{uri: "http://"+this.state.ipserver+"/assets/images/cewek.png"}} />;
      const cowok = <Image style={styles.pic} source={{uri: "http://"+this.state.ipserver+"/assets/images/cowok.png"}} />;
      const photo = <Image style={styles.pic} source={{uri: "http://"+this.state.ipserver+"/storage/guru/"+data.foto}} />;      
      if(data.foto){return photo;}
      else if(data.jenis === "Perempuan"){return cewek;}
      else{return cowok;}
    }    

    searchUpdated(term) {
      this.setState({ searchTerm: term })
    }

    onRefresh() {
      this.setState({ users: [] });      
      //Call the Service to get the latest data
      this.fetchGuru();
    }

    fetchGuru() {
      axios.get("http://"+this.state.ipserver+"/api/alluser")
          .then(response => {
              const data = response.data.data;
              var datafilter = data.filter(e => e.id !== this.state.user.id)
              this.setState({ users: datafilter });
              this.setState({ refreshing: false });
          }).catch(error => {           
            
              if(error.message === "Network Error"){            
                  this.refs.toast.show('Maaf Jaringan Tidak Tersedia !');                  
              }
              if(error.response.status == 401){
                  this.refs.toast.show(error.response.data.message);                  
              }
              
          });
    };
    // -----    
}

const styles = StyleSheet.create({
  DataItem:{
    borderBottomWidth: 0.5,
    borderColor: '#e6e6e6',
    padding: 10,
    flexDirection:'row',
  },
  search:{
    color:'#000',
    top:12,
    left:9,
    marginRight:7,
  },
  caribox:{
    flex:1,  
    marginLeft:8,
    marginRight:8,
    alignSelf: 'center',
    alignItems: 'stretch',
    justifyContent: 'space-between',               
    textAlign: 'right',
  },
  header:{
    flexDirection:'row',
    borderColor: '#e6e6e6',
    borderWidth: 1,
    marginTop:0,   
    height:55,  
    backgroundColor:'white',
    opacity: 0.9,
    elevation:3,   
  },
  mapel: {
    color: 'rgba(0,0,0,0.5)',
    fontSize:12,
  },
  konten:{
    marginLeft:5,
  },
  searchInput:{
    fontWeight:'bold',
    padding: 6,    
    height:35,      
    backgroundColor:'#fefefe',
  },
  background: {           
    resizeMode: 'stretch',
    height:55,
    elevation:1,
  },
  pic:{        
    width:25,
    height:38,            
  },
  container: {
    flex: 1,        
    justifyContent: 'flex-start',
    backgroundColor:'#fff',  
  },
});