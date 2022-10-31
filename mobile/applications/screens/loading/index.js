import * as React from 'react';
import { View , StyleSheet , Text } from 'react-native';
import { MaterialIndicator } from 'react-native-indicators';
import MyStatusBar from '../../components/statusBar';

export default class LoadingScreen extends React.Component {

constructor(props) {
    super(props);    
}

render() {
	return (
	<View style={styles.container}>
      <MyStatusBar backgroundColor={global.BarColor} barStyle="dark-content" />         
      <Text style={styles.textLoading}>Memuat ...</Text>
	</View>
	);
}
//function

//end
}

const styles = StyleSheet.create({  
  container: {    
    backgroundColor: "#f3f3f3",    
    alignItems:"center",    
    justifyContent:"center",
    flex:1
  },
  textLoading : {
    marginTop:10,
    fontSize:20,
  },
});