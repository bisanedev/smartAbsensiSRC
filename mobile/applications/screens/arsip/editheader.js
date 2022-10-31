import React from 'react';
import {      
    StyleSheet,
    View,
    Text,
    TouchableOpacity,  
    ImageBackground,
} from 'react-native';
import Icon from 'react-native-vector-icons/Ionicons';

export default class Header extends React.Component {
    constructor(props) {
        super(props);          
    }

    render() {
        let data = this.props.tanggal;        
        return(
            <View style={styles.container}>   
            <ImageBackground source={require('../../assets/pattern.jpg')} style={styles.background}>
            <View style={styles.navbar}>                                           
               <TouchableOpacity onPress={this.TombolBack} style={styles.tombol}>
                    <Icon name="ios-arrow-back" color="#000" size={30}/>
                    <Text style={styles.backText}>KEMBALI</Text> 
               </TouchableOpacity>                                                      
               <View style={styles.title}>
                    <Text style={styles.titleText}>Edit Absensi Siswa</Text>
               </View>               
            </View>                
            </ImageBackground>
            </View>
        );
    }
    //function    
    TombolBack = () => {          
        this.props.kembali();
    };
    TombolPiket = () => {
        this.props.piket();
    };

}

const styles = StyleSheet.create({
    container:{   
        elevation:1,              
    },   
    title:{    
        flex:1,        
        alignSelf: 'center',
        alignItems: 'center',
        justifyContent: 'space-between',               
        textAlign: 'right',
    },
    titleText:{                
        fontSize:16,
        fontWeight:'bold',
        textTransform: 'uppercase',
    },   
    navbar:{
        borderColor: '#e6e6e6',
        borderWidth: 1,
        marginTop:0,
        height:55,        
        elevation:3,
        paddingHorizontal:5,
        backgroundColor:'white',
        opacity: 0.9,
        flex:1,   
        flexDirection:'row',
        alignContent:'center',        
        justifyContent:'space-between',
        shadowOffset:{  width: 10,  height: 10,  },        
    },
    tombol:{           
        flexDirection:'row',     
        alignSelf: 'center',
        alignItems: 'center',
        justifyContent: 'space-between',
    },
    backText:{
        marginLeft:7,
        fontSize:11,
        fontWeight:'bold',
    },  
    background: {           
        resizeMode: 'stretch',
        height:55,        
    },
});