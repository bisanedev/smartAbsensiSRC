import React from 'react';
import { StyleSheet,View,Image,Text,TextInput,TouchableOpacity } from 'react-native';
import AsyncStorage from '@react-native-async-storage/async-storage';
import {widthPercentageToDP as wp, heightPercentageToDP as hp} from 'react-native-responsive-screen';
import RBSheet from "react-native-raw-bottom-sheet";

export default class Body extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
        ipaddress: '',
        storevar: {            
            id:this.props.data.siswa_id,            
            kode:this.props.data.kode,
            kodenote:this.props.data.kode_note,
          }
        };

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
        const photo = <Image style={styles.pic} source={{uri: "http://"+this.state.ipaddress+"/storage/siswa/"+data.foto}} />;
        let profile;
        if(data.foto){ profile = photo }
        else if(data.jenis === "Perempuan"){profile = cewek}else{profile = cowok}
        //--- note
        let note;        
        if(this.state.storevar.kode == "s" || this.state.storevar.kode == "i"){
            note=
            <View>
                <TextInput
                    style={styles.note}
                    value={this.state.storevar.kodenote}
                    onChangeText={(ket) => this.noteChange(ket)}                    
                    placeholder='Input Keterangan'
                />
            </View>            
        }
        //--- render template
        return (
            <View style={styles.container}>
                <View style={styles.card}>
                <Text style={styles.urut}>{data.urut}</Text> 
                {profile}                 
                <View style={styles.konten}>                    
                    <View style={styles.infosiswa}>
                        <Text style={styles.nama}>
                        { ((data.nama).length > 32) ? (((data.nama).substring(0,32-3)) + '...') : data.nama }
                        </Text>
                        <Text style={styles.jenis}>{data.jenis}</Text>                                               
                    </View>
                </View>
                <TouchableOpacity style={this.menuStyle(this.state.storevar.kode)} onPress={() => this.RBSheet.open()}>                                                           
                    <Text style={{fontSize:11,fontWeight:'bold',color:'white'}}>{this.absenText(this.state.storevar.kode)}</Text>                   
                    <RBSheet ref={ref => { this.RBSheet = ref; }} height={hp("43%")} openDuration={250} customStyles={{container: {padding:12}}}>
                        <Text style={{fontWeight:"bold",color:"#757575",marginBottom:10}}>{ ((data.nama).length > 32) ? (((data.nama).substring(0,32-3)) + '...') : data.nama }</Text>
                        <TouchableOpacity style={{backgroundColor:"green",padding:8,marginBottom:5}} onPress={() => {this.pressHandler('h');this.RBSheet.close();}}>
                            <Text style={{color:"#fff",textAlign:"center"}}>Hadir</Text>
                        </TouchableOpacity>    
                        <TouchableOpacity style={{backgroundColor:"#07ade3",padding:8,marginBottom:5}} onPress={() => {this.pressHandler('s');this.RBSheet.close();}}>
                            <Text style={{color:"#fff",textAlign:"center"}}>Sakit</Text>
                        </TouchableOpacity>
                        <TouchableOpacity style={{backgroundColor:"blue",padding:8,marginBottom:5}} onPress={() => {this.pressHandler('i');this.RBSheet.close();}}>
                            <Text style={{color:"#fff",textAlign:"center"}}>Izin</Text>
                        </TouchableOpacity>
                        <TouchableOpacity style={{backgroundColor:"orange",padding:8,marginBottom:5}} onPress={() => {this.pressHandler('t');this.RBSheet.close();}}>
                            <Text style={{color:"#fff",textAlign:"center"}}>Terlambat</Text>
                        </TouchableOpacity>
                        <TouchableOpacity style={{backgroundColor:"red",padding:8}} onPress={() => {this.pressHandler('a');this.RBSheet.close();}}>
                            <Text style={{color:"#fff",textAlign:"center"}}>Alpa</Text>
                        </TouchableOpacity>                    
                    </RBSheet>
                </TouchableOpacity>                      
                </View>
                {note}
            </View>            
        );
    }        
    //-----   
    absenText = (code) => {
        let absen;        
        if(code === 'h'){absen = "Hadir" }
        if(code === 's'){absen = "Sakit" }
        if(code === 'i'){absen = "Izin" }
        if(code === 't'){absen = "Terlambat" }
        if(code === 'a'){absen = "Alpa" }
        return absen;
    }

    noteChange = async (value) => {
        await this.setState(prevState => {
            let storevar = Object.assign({}, prevState.storevar);
            storevar.kodenote = value;
            return { storevar };
        })        
        this.props.absen(this.state.storevar.id,this.state.storevar.kode,this.state.storevar.kodenote);
    }

    pressHandler = async (value) => {                
        await this.setState(prevState => {
            let storevar = Object.assign({}, prevState.storevar);
            storevar.kode = value;
            storevar.kodenote = "";
            return { storevar };
        })    
        this.props.absen(this.state.storevar.id,this.state.storevar.kode,this.state.storevar.kodenote);     
    }

    menuStyle = function(code) {
        let colors;        
        if(code === 'h'){colors = "green" }
        if(code === 's'){colors = "#07ade3" }
        if(code === 'i'){colors = "blue" }
        if(code === 't'){colors = "orange" }
        if(code === 'a'){colors = "red" }
        return {          
          backgroundColor:colors,
          width:70,
          height:70,
          alignItems:"center",
          justifyContent:"center"
        }
    }  
//----    
}

const styles = StyleSheet.create({
    container: {
      flex: 1,          
    },
    konten:{              
        alignSelf: 'center',
        alignItems: 'center',               
    },
    urut:{
        position:'absolute',        
        backgroundColor:'white',
        bottom:0,
        left:0,
        zIndex:1,
        paddingLeft:2,
        paddingRight:2,
        fontSize:10,        
        fontWeight:'bold',        
    },  
    pic:{        
        width:50,
        height:70,        
    },
    infosiswa:{        
        textAlign:'center',
        alignContent:'center',        
    },
    nama:{        
        textTransform:'capitalize',
        fontSize:14,        
        textAlign:'center',
        fontWeight:'bold',
    },
    jenis:{
        fontSize:10,
        textAlign:'center',        
    },
    card:{
        flex:1,
        backgroundColor:'#f9f9f9',      
        flexDirection:'row', 
        justifyContent:'space-between', 
        alignItems:'center',                 
        marginTop:15,                  
        marginLeft:10,
        marginRight:10,
        borderColor: '#e6e6e6',
        borderWidth: 0.7, 
        borderTopRightRadius:11,
        borderBottomRightRadius:11,
    },
    note:{
        height: 50,        
        backgroundColor:'#f9f9f9',
        borderBottomWidth:0.7,
        borderBottomColor:'#e6e6e6',
        borderLeftWidth:0.7,
        borderLeftColor:'#e6e6e6',
        borderRightWidth:0.7,
        borderRightColor:'#e6e6e6',
        padding:11,
        marginLeft:10,
        marginRight:10, 
    },
});
