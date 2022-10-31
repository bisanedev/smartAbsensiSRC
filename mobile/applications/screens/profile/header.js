import React from 'react';
import {      
    StyleSheet,
    View,
    Text,    
    ImageBackground,
} from 'react-native';

export default class Header extends React.Component {
    constructor(props) {
        super(props);          
    }

    render() {        
        return(
            <View style={styles.container}>   
            <ImageBackground source={require('../../assets/pattern.jpg')} style={styles.background}>
            <View style={styles.navbar}>                                                                                                               
               <View style={styles.title}>
                    <Text style={styles.titleText}>Profil Guru</Text>
               </View>               
            </View>                
            </ImageBackground>
            </View>
        );
    }
}

const styles = StyleSheet.create({
    container:{   
        elevation:1,              
    },
    piketText:{
        marginRight:7,
        fontSize:11,
        fontWeight:'bold',        
    },
    title:{
        alignSelf: 'center',
        alignItems: 'center',                   
    },
    titleText:{
        fontSize:19,
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
        justifyContent: 'center',        
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