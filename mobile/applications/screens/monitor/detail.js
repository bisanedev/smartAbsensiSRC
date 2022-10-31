import React from 'react';
import {      
  StyleSheet,
  View,
  Text,    
  ScrollView,  
} from 'react-native';
import MyStatusBar from '../../components/statusBar';
import Header from './headerDetails';
import Body from './body';

export default class DetailScreen extends React.Component { 
    constructor(props) {
        super(props);
        this.state = {};
        
        this.data = this.props.route.params.monitor; 
        this.title = this.props.route.params.title;
    }

    static navigationOptions =  {
        headerShown: false,   
    };    
    
    render() {      
      //------------------ 
      return (
        <View style={styles.container}>
            <MyStatusBar backgroundColor="#fff" barStyle="dark-content" />
            <Header kembali={this._goBack} title={this.title}/>
            <ScrollView>
            { this.data.sort((a, b) => a.nama > b.nama).map((item,i) => (
              <Body data={item} key={i}/>
            ))}                
            </ScrollView>            
        </View>
      );
    }
    _goBack = () => {            
        this.props.navigation.goBack();
    };
}

const styles = StyleSheet.create({
    container: {
      flex: 1, 
      backgroundColor:'#f2f1f1',       
    }
});