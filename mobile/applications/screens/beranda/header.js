import React from 'react';
import {      
    StyleSheet,
    View,
    Text,
    Image,  
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
               <Image source={{uri:this.props.logo}} style={styles.logo}/>
               <View style={styles.title}>
                    <Text style={styles.titleText}>{this.props.nama}</Text>
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
    logo:{
        width:44,
        height:44,                
        flexDirection:'row',     
        alignSelf: 'center',
        alignItems: 'center',
        justifyContent: 'space-between'             
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
    background: {           
        resizeMode: 'stretch',
        height:55,        
    },
});