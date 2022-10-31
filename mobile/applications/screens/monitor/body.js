import React from 'react';
import { StyleSheet,View,Image,Text } from 'react-native';
import { withNavigation } from '@react-navigation/compat';
import AsyncStorage from '@react-native-async-storage/async-storage';
import moment from 'moment';

class Body extends React.Component {
    constructor() {
        super();     
        this.state = {ipaddress: ''};
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
        let keterangan;
        if(data.kode_note){keterangan = <Text style={styles.keterangan}> "{data.kode_note}" </Text>}
        ///--------------
        return (
            <View style={styles.container}>
                <View style={styles.items}>
                {profile}
                <View style={styles.konten}>
                    <Text style={styles.txtnama}>{data.nama.toLowerCase()}</Text>                                        
                    <Text style={styles.txt}>Kelas [{data.kelas}] ~ Mata Pelajaran {data.mapel}</Text>
                    <Text style={styles.txt}> {data.mapel_guru}</Text>
                    <Text style={styles.txt}>Jam {moment(data.mapel_start).format("HH:mm")} ~ {moment(data.mapel_end).format("HH:mm")}</Text>
                    {keterangan}
                </View>                                                                                
                </View>
            </View>
        );
    }    

}

const styles = StyleSheet.create({
    container: {
      flex: 1,          
    },
    items:{
        backgroundColor:"#fff",
        flexDirection:'row', 
        justifyContent:'space-between', 
        alignItems:'center',                 
        marginTop:10,
        marginBottom:10,          
        marginLeft:10,
        marginRight:10,
        borderColor: '#757575',
        borderWidth: 0.4,          
    },
    pic:{        
        width:80,
        height:120,        
    },
    konten:{        
        marginLeft:-20,
        alignItems: 'center',        
        flex: 1,
    },
    txt:{
        marginTop:3,
        fontSize:11,
        color:'#757575',
    },   
    keterangan:{
        marginTop:3,    
        fontSize:10,
        fontWeight:'bold',
        color:'#757575',
    },
    txtnama:{  
        textTransform: 'capitalize',      
        fontWeight:'bold',
        fontSize:16,
        marginBottom:5,        
        color:'#000',
    }

});

export default withNavigation(Body);