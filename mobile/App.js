import * as React from 'react';
import { Text, View } from 'react-native';
import { NavigationContainer } from '@react-navigation/native';
import { createStackNavigator } from '@react-navigation/stack';
import { createBottomTabNavigator } from '@react-navigation/bottom-tabs';
import AsyncStorage from '@react-native-async-storage/async-storage';
import SplashScreen from 'react-native-splash-screen';
import Icon from 'react-native-vector-icons/MaterialIcons';
import axios from 'axios';
import LoadingScreen from './applications/screens/loading';
import SignInScreen from './applications/screens/authentikasi';
import BerandaScreen from './applications/screens/beranda';
import AbsenScreen from './applications/screens/absensi';
import JadwalScreen from './applications/screens/jadwal';
import GuruScreen from './applications/screens/jadwal/guru';
import MonitorScreen from './applications/screens/monitor';
import DetailScreen from './applications/screens/monitor/detail';
import ArsipScreen from './applications/screens/arsip';
import ArsipDetailScreen from './applications/screens/arsip/detail';
import ArsipEditScreen from './applications/screens/arsip/edit';
import ProfileScreen from './applications/screens/profile';

const Stack = createStackNavigator();
const AuthContext = React.createContext();
const Tab = createBottomTabNavigator();
global.BarColor = '#f3f3f3';

export default function App() {
    
  const [state, dispatch] = React.useReducer(
    (prevState, action) => {
      switch (action.type) {
        case 'RESTORE_TOKEN':
          return {
            ...prevState,
            userToken: action.token,                      
            isLoading: false,
          };
        case 'SIGN_IN':
          return {
            ...prevState,
            isSignout: false,
            userToken: action.token,            
          };
        case 'SIGN_OUT':
          return {
            ...prevState,
            isSignout: true,
            userToken: null,            
          };
      }
    },
    {
      isLoading: true,
      isSignout: false,
      userToken: null,      
    }
  );
  
  //didmounted
  React.useEffect(() => {   
    const bootstrapAsync = async () => {
      let userToken;        
      try {
        userToken = await AsyncStorage.getItem('userToken');             
      } catch (e) {
        // Restoring token failed
      }
      // After restoring token, we may need to validate it in production apps
      SplashScreen.hide(); // hide splashscreen     
      axios.defaults.headers.common['Authorization'] = userToken;                
      dispatch({ type: 'RESTORE_TOKEN', token: userToken });      
    };    
    bootstrapAsync();    
  }, []);    

  const authContext = React.useMemo(
    () => ({
      signIn: (data) => {        
        dispatch({ type: 'SIGN_IN', token: data});        
      },    
      signOut: async () => {
        await AsyncStorage.removeItem('userToken');                
        dispatch({ type: 'SIGN_OUT' });
      }    
    }),
    []
  );    

  return (
    <AuthContext.Provider value={authContext}>
      <NavigationContainer>
        <Stack.Navigator>
          {state.isLoading ? (
            // We haven't finished checking for the token yet
            <Stack.Screen name="Splash" component={LoadingScreen} options={{headerShown: false}} />
          ) : state.userToken == null ? (
            // No token found, user isn't signed in            
            <Stack.Screen name="SignIn" options={{headerShown: false, animationTypeForReplace: state.isSignout ? 'pop' : 'push',}}>
            {props => <SignInScreen {...props} login={authContext.signIn} />}
            </Stack.Screen>
          ) : (
            // protected page
            <Stack.Screen name="MenuTabs" component={MenuTabs} options={{headerShown: false}}/>             
            // end protected page
          )}
        </Stack.Navigator>
      </NavigationContainer>
    </AuthContext.Provider>
  );
}

function MenuTabs() {
  const { signOut } = React.useContext(AuthContext);
  return (
    <Tab.Navigator
      screenOptions={({ route }) => ({
        tabBarIcon: ({ focused, color, size }) => {
          let iconName;
          
          if (route.name === 'home') {
            iconName = 'schedule'
          } else if (route.name === 'jadwal') {
            iconName = 'event'
          } else if (route.name === 'monitor') {
            iconName = 'timeline'
          } else if (route.name === 'arsip') {
            iconName = 'folder-open'
          } else if (route.name === 'profile') {
            iconName = 'account-box'
          }

          // You can return any component that you like here!
          return <Icon name={iconName} size={size} color={color} />;
        }   
      })}
      tabBarOptions={{
        initialRouteName: 'home',
        activeTintColor : '#000',        
        inactiveTintColor : '#757575',
        backgroundColor: '#f9f9f9',
        labelStyle: {
          fontSize: 10,
          fontWeight:'bold',
        },
      }}
    >
      <Tab.Screen name="home" component={HomeStackGroup} options={{tabBarLabel: 'Jadwal Saya'}}/>
      <Tab.Screen name="jadwal" component={JadwalStackGroup} options={{tabBarLabel: 'Jadwal'}}/>
      <Tab.Screen name="monitor" component={MonitorStackGroup} options={{tabBarLabel: 'Monitoring'}}/>
      <Tab.Screen name="arsip" component={ArsipStackGroup} options={{tabBarLabel: 'Arsip'}}/>
      <Tab.Screen name="profile" options={{tabBarLabel: 'Profil'}}>
        {props => <ProfileScreen {...props} logout={signOut} />}
      </Tab.Screen>
    </Tab.Navigator>
  );
}

const HomeStack = createStackNavigator();
function HomeStackGroup() {
  return (    
    <HomeStack.Navigator initialRouteName="Home">
      <HomeStack.Screen name="Home" component={BerandaScreen} options={{headerShown: false}}/>
      <HomeStack.Screen name="Absen" component={AbsenScreen} options={{headerShown: false}}/>
    </HomeStack.Navigator>
  );
}

const JadwalStack = createStackNavigator();
function JadwalStackGroup() {
  return (  
    <JadwalStack.Navigator initialRouteName="Jadwal">
        <JadwalStack.Screen name="Jadwal" component={JadwalScreen} options={{headerShown: false}}/> 
        <JadwalStack.Screen name="Guru" component={GuruScreen} options={{headerShown: false}}/>
        <JadwalStack.Screen name="AbsenJadwal" component={AbsenScreen} options={{headerShown: false}}/>                
    </JadwalStack.Navigator>
  );
}

const MonitorStack = createStackNavigator();
function MonitorStackGroup() {
  return (  
    <MonitorStack.Navigator initialRouteName="Monitor">
        <MonitorStack.Screen name="Monitor" component={MonitorScreen} options={{headerShown: false}}/> 
        <MonitorStack.Screen name="Detail" component={DetailScreen} options={{headerShown: false}}/> 
    </MonitorStack.Navigator>
  );
}

const ArsipStack = createStackNavigator();
function ArsipStackGroup() {
  return (  
    <ArsipStack.Navigator initialRouteName="Arsip">
        <ArsipStack.Screen name="Arsip" component={ArsipScreen} options={{headerShown: false}}/> 
        <ArsipStack.Screen name="Detail" component={ArsipDetailScreen} options={{headerShown: false}}/> 
        <ArsipStack.Screen name="Edit" component={ArsipEditScreen} options={{headerShown: false}}/>         
    </ArsipStack.Navigator>
  );
}