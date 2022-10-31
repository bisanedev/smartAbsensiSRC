import Vue from 'vue';
import VueRouter from 'vue-router';

Vue.use(VueRouter)

import HomePage from './layouts/dashboard.vue';
import LoginPage from './layouts/login.vue';
import NotFoundPage from './layouts/notfound.vue';

const routes = [
{
  path: '/login',
  name: 'login',
  component: LoginPage
},
{ path: '/', redirect: { name: 'dashboard' } },
{
  path: '/',  
  component: HomePage,
  children:[   
    {
      path: 'dashboard',      
      component: require('./pages/dashboard/index.vue').default,
      name: 'dashboard',
    },
    {
      path: 'jadwal/kelas',      
      component: require('./pages/jadwal/jadwalkelas.vue').default,
      name: 'jadwalkelas',
      meta: { akses: "admin" }, 
    },
    {
      path: 'jadwal/guru',      
      component: require('./pages/jadwal/jadwalguru.vue').default,
      name: 'jadwalguru',
      meta: { akses: "admin" }, 
    },
    {
      path: 'absensi/mapel',      
      component: require('./pages/absensi/index.vue').default,
      name: 'absensi',
      meta: { akses: "admin" }, 
    },
    {
      path: 'absensi/siswa',      
      component: require('./pages/absensi/siswa.vue').default,
      name: 'absensisiswa',
      meta: { akses: "admin" }, 
    },
    {
      path: 'input/guru',      
      component: require('./pages/guru/guru.vue').default,
      name: 'guru',
      meta: { akses: "admin" },  
    },
    {
      path: 'input/kelas',      
      component: require('./pages/kelas/kelas.vue').default,
      name: 'kelas',
      meta: { akses: "admin" },  
    },
    {
      path: 'input/siswa',      
      component: require('./pages/siswa/siswa.vue').default,
      name: 'siswa',
      meta: { akses: "admin" },  
    },
    {
      path: 'input/semester',      
      component: require('./pages/semester/index.vue').default,
      name: 'semester',
      meta: { akses: "admin" },  
    },
    {
      path: 'arsip/guru',      
      component: require('./pages/guru/arsip.vue').default,
      name: 'arsipguru',
      meta: { akses: "admin" },  
    },
    {
      path: 'arsip/siswa',      
      component: require('./pages/siswa/arsip.vue').default,
      name: 'arsipsiswa',
      meta: { akses: "admin" },  
    },
    {
      path: 'sistem',      
      component: require('./pages/sistem/index.vue').default,
      name: 'sistem',
      meta: { akses: "admin" },  
    },
    {
      path: 'rekap/kelas',      
      component: require('./pages/rekap/kelas/index.vue').default,
      name: 'rekapkelas',
      meta: { akses: "admin" },      
    },
    {
      path: 'rekap/guru',      
      component: require('./pages/rekap/guru/index.vue').default,
      name: 'rekapguru',
      meta: { akses: "admin" },    
    },
    {
      path: 'rekap/siswa',      
      component: require('./pages/rekap/siswa/index.vue').default,
      name: 'rekapsiswa',
      meta: { akses: "admin" },   
    },
    {
      path: 'laporan/absensi',      
      component: require('./pages/laporan/absensi/index.vue').default,
      name: 'laporanabsensi',      
    },
    {
      path: 'laporan/clock',      
      component: require('./pages/laporan/clock/index.vue').default,
      name: 'laporanclock',      
    },
    {
      path: 'laporan/piket',      
      component: require('./pages/laporan/piket/index.vue').default,
      name: 'laporanpiket',      
    }
  ]
},
{ path: "*", component: NotFoundPage  }
]

const router = new VueRouter({ routes });

router.beforeEach((to, from, next) => {
  // redirect to login page if not logged in and trying to access a restricted page
  const publicPages = ['/login'];
  const authRequired = publicPages.includes(to.path);
  const loggedIn = window.localStorage.getItem('token');    
  // login
  if (!authRequired && !loggedIn) {
    console.log('not authorized')
    router.push({ path: 'login' })    
  }else if (authRequired && loggedIn){
    console.log('authorized')
    router.push({ path: 'dashboard' })     
  }else{
    //meta akses
    if (!to.meta.akses){
      next() 
    }else{  
      var akses = jwt_decode(loggedIn);              
      if (akses.role == to.meta.akses || akses.role == "superuser" ) {
        console.log('authorized')      
        next()
      } else {
       Vue.notify({
        group: 'notifikasi',                  
        title: 'Not Authorized ',
        type: 'error',
        text: 'Butuh Role '+to.meta.akses+' Akses Untuk Mengakses'
      });
        console.log('not authorized')
        router.push({ path: 'dashboard' })
      }
    }
    //end next()
  }  
})


export default router