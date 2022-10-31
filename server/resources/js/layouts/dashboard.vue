<template>
    <div class="page">
        <div class="page-main">
            <div class="header py-4">
                <div class="container">
                    <div class="d-flex">
                        <a class="header-brand" href="./">
                            <img src="storage/logo.png" class="header-brand-img" alt="tabler logo">
                            {{sekolah}}
                        </a>
                        <div class="d-flex order-lg-2 ml-auto">
                            <div class="dropdown">
                                <a href="#" class="nav-link pr-0 leading-none" data-toggle="dropdown">
                                    <span v-if="data.foto" class="avatar" v-bind:style="{'background-image':'url(./storage/guru/'+ data.foto +')' }">        
                                    </span>
                                    <span v-else-if="data.jenis == 'Perempuan'" class="avatar" style="background-image: url(./assets/images/cewek.png)">
                                    </span>
                                    <span v-else class="avatar" style="background-image: url(./assets/images/cowok.png)">                                     
                                    </span>
                                    <span class="ml-2 d-none d-lg-block">
                                        <span class="text-default">{{data.username}}</span>
                                        <small class="text-muted d-block mt-1" style="text-transform: capitalize;">{{data.name}}</small>
                                    </span>
                                </a>
                                <div class="dropdown-menu dropdown-menu-right dropdown-menu-arrow">
                                    <router-link to="/sistem" class="dropdown-item" active-class="active">
                                        <i class="dropdown-icon fe fe-database"></i> Sistem
                                    </router-link>
                                    <div class="dropdown-divider"></div>
                                    <button class="dropdown-item" :disabled="disable" v-on:click="logout()">
                                        <i class="dropdown-icon fe fe-log-out"></i> Sign out
                                    </button>
                                </div>
                            </div>
                        </div>
                        <a href="#" class="header-toggler d-lg-none ml-3 ml-lg-0" data-toggle="collapse" data-target="#headerMenuCollapse">
                            <span class="header-toggler-icon"></span>
                        </a>
                    </div>
                </div>
            </div>
            <div class="header collapse d-lg-flex p-0" id="headerMenuCollapse">
                <div class="container">
                    <div class="row align-items-center">
                        <div class="col-lg-3 ml-auto">
                            <div class="float-right">IP Server <b>{{ipaddress}}</b></div>
                        </div>
                        <div class="col-lg order-lg-first">
                            <ul class="nav nav-tabs border-0 flex-column flex-lg-row">
                                <li class="nav-item">
                                    <router-link to="/dashboard" class="nav-link" active-class="active">
                                        <i class="fe fe-home"></i>
                                        Dashboard
                                    </router-link>
                                </li>
                                <li class="nav-item dropdown">
                                    <a href="javascript:void(0)" class="nav-link" data-toggle="dropdown">
                                        <i class="fe fe-calendar"></i> Jadwal
                                    </a>
                                    <div class="dropdown-menu dropdown-menu-arrow">
                                        <router-link to="/jadwal/kelas" class="dropdown-item" active-class="active">Kelas</router-link>
                                        <router-link to="/jadwal/guru" class="dropdown-item" active-class="active">Guru</router-link>
                                    </div>
                                </li>
                                <li class="nav-item dropdown">
                                     <a href="javascript:void(0)" class="nav-link" data-toggle="dropdown">
                                        <i class="fe fe-clipboard"></i> Absensi
                                    </a>
                                    <div class="dropdown-menu dropdown-menu-arrow">                             
                                        <router-link to="/absensi/mapel" class="dropdown-item" active-class="active">Mapel</router-link>  
                                        <router-link to="/absensi/siswa" class="dropdown-item" active-class="active">Siswa</router-link>                             
                                    </div>                                    
                                </li>
                                <li class="nav-item dropdown">
                                    <a href="javascript:void(0)" class="nav-link" data-toggle="dropdown">
                                        <i class="fe fe-edit-3"></i> Input Data
                                    </a>
                                    <div class="dropdown-menu dropdown-menu-arrow">
                                        <router-link to="/input/guru" class="dropdown-item" 
                                        active-class="active">Guru</router-link>
                                        <router-link to="/input/kelas" class="dropdown-item" 
                                        active-class="active">Kelas</router-link>
                                        <router-link to="/input/siswa" class="dropdown-item" 
                                        active-class="active">Siswa</router-link>  
                                        <router-link to="/input/semester" class="dropdown-item" 
                                        active-class="active">Semester</router-link>          
                                    </div>
                                </li>
                                <li class="nav-item dropdown">
                                    <a href="javascript:void(0)" class="nav-link" data-toggle="dropdown">
                                        <i class="fe fe-folder"></i> Arsip
                                    </a>
                                    <div class="dropdown-menu dropdown-menu-arrow">
                                        <router-link to="/arsip/guru" class="dropdown-item" active-class="active">Data Guru</router-link>                                        
                                        <router-link to="/arsip/siswa" class="dropdown-item" active-class="active">Data Siswa</router-link>
                                    </div>
                                </li>
                                <li class="nav-item dropdown">
                                    <a href="javascript:void(0)" class="nav-link" data-toggle="dropdown">
                                        <i class="fe fe-pie-chart"></i> Rekap
                                    </a>
                                    <div class="dropdown-menu dropdown-menu-arrow">
                                      <router-link to="/rekap/guru" class="dropdown-item" active-class="active">Guru</router-link>
                                      <router-link to="/rekap/kelas" class="dropdown-item" active-class="active">Kelas</router-link>                                      
                                      <router-link to="/rekap/siswa" class="dropdown-item" active-class="active">Siswa</router-link>
                                    </div>            
                                </li>
                                <li class="nav-item dropdown">
                                    <a href="javascript:void(0)" class="nav-link" data-toggle="dropdown">
                                        <i class="fe fe-printer"></i> Laporan
                                    </a>
                                    <div class="dropdown-menu dropdown-menu-arrow">
                                        <router-link to="/laporan/absensi" class="dropdown-item" active-class="active">Absensi</router-link>
                                        <router-link to="/laporan/clock" class="dropdown-item" active-class="active">Clock</router-link>
                                        <router-link to="/laporan/piket" class="dropdown-item" active-class="active">Clock Piket</router-link>
                                    </div>            
                                </li>
                            </ul>
                        </div>
                    </div>
                </div>
            </div>
            <div class="my-3 my-md-5">
                <router-view></router-view>
            </div>
        </div>
        <v-dialog />
    </div>
</template>

<script type="application/javascript">

export default {
    data() {
        return {
            disable: false,
            data: [],
            ipaddress: '', 
            sekolah:'',
        }
    },
    created() {
        this.RenderTokenData();
        this.RenderWebsocket();
    },
    mounted() {        
        if (localStorage.ipaddress) {
            this.ipaddress = localStorage.ipaddress;
            this.sekolah = localStorage.sekolah;
            document.documentElement.style.setProperty('--mycolor', localStorage.warna);
        }            
    },
    methods: {
        logout: function() {
            //window.localStorage.removeItem('token');
            this.disable = true
            window.localStorage.clear();
            delete window.axios.defaults.headers.common['Authorization'];
            this.$router.replace({ name: 'login' });

        },
        RenderWebsocket: function() {
          this.socket = new WebSocket("ws://"+window.location.host+"/websocket/absen/ws");
          this.socket.onopen = () => {
            console.log("websocket connected");                     
            this.socket.onmessage = ({data}) => {          
              var socketData=JSON.parse(data);
              
              this.$notify({
                group: 'notifikasi',                         
                title: socketData.Data,
                type: 'info',
                text: 'Guru Masuk Kelas & Memulai Pelajaran'
              });              
            };
          }

          this.socket.onerror = () =>  {          
              console.log("error");
          }

          this.socket.onclose = () => {        
            console.log("disconnected");     
          }
        },
        RenderTokenData: function() {
            var akses = jwt_decode(window.localStorage.getItem('token'));
            axios.get(window.location.origin + '/api/users/' + akses.id)
                .then(response => {
                    if (response.data.status == "success") {
                        this.data = response.data.data
                    }
                })
                .catch(error => {
                    if (error.response.status === 401) {
                        this.$notify({
                            group: 'notifikasi',
                            title: 'Render Token',
                            type: 'alert',
                            text: error.response.data.message
                        });
                    }
                    //console.log(error);
                })
            //------
        }
        //---
    }
    // ...
}
</script>