<template>
  <div class="container">
  	<div class="page-header">
       <h1 class="page-title">Sistem Aplikasi</h1>
    </div>
    <div class="row row-cards">
      <div class="col-6 col-lg-6">
        <div class="card">
        <div class="card-header">
          <h3 class="card-title">Password Piket</h3>
          <div class="card-options">
             <button @click="gantiPasswordPiket()" class="btn btn-primary"><i class="fe fe-lock"></i> Ganti Password</button>
          </div>
        </div>
        <div class="card-body">
          <input type="text" @keydown.space.prevent class="form-control" style="text-align:center;height: 80px;font-size: xx-large;background-color:grey;font-weight: bold;" placeholder="Password Piket" v-model="password">          
        </div>
        </div>
      </div>
      <div class="col-6 col-lg-6">
        <div class="card">
        <div class="card-header">
          <h3 class="card-title">Password Laporan Clock [Excel]</h3>
          <div class="card-options">
             <button @click="gantiPasswordClock()" class="btn btn-primary"><i class="fe fe-lock"></i> Ganti Password</button>
          </div>
        </div>
        <div class="card-body">
          <input type="text" @keydown.space.prevent class="form-control" style="text-align:center;height: 80px;font-size: xx-large;background-color:grey;font-weight: bold;" placeholder="Password File Excel" :value="passwordExcel.toLowerCase()" @input="passwordExcel = $event.target.value.toLowerCase()">          
        </div>
        </div>
      </div>
      <div class="col-6 col-lg-6">
        <div class="card">
        <div class="card-header">
          <h3 class="card-title">Backup Database</h3>
          <div class="card-options">
             <button @click="BackupDatabase()" class="btn btn-primary"><i class="fe fe-refresh-cw"></i> Generate New Backup</button>
          </div>
        </div>
        <div class="card-body">
          <input type="text" @keydown.space.prevent class="form-control" style="text-align:center;height: 80px;font-size: xx-large;background-color:grey;font-weight: bold;" placeholder="File Backup" v-model="namaDB" disabled="disabled">          
        </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script type="application/javascript">

import moment from "moment";

export default { 
  components: {
    
  },
  data () {
    return {
      moment:moment,
      password:'',
      namaDB:'',
      passwordExcel:'',
    }
  },
  created() {
   this.fecthPassword();
  },
  methods: {
    /** Method **/
    fecthPassword(){
        axios.get(window.location.origin + "/api/sistem/password")
        .then(response => {
          this.password = response.data.data.password_piket
          this.passwordExcel = response.data.data.password_clock
          this.namaDB = response.data.data.backup_database
        }).catch(error => this.password = this.passwordExcel = "")
    },
    BackupDatabase(){
      axios.get(window.location.origin + "/api/sistem/backup/export")
        .then(response => {
          this.namaDB = response.data.file          
      }).catch(error => this.namaDB = "")
    },
    gantiPasswordClock(){
      if(this.passwordExcel == ''){
        this.$notify({
              group: 'notifikasi',                  
              title: 'Update Data Gagal',
              type: 'error',
              text: "Periksa Kembali Input Data"
        }); 
        return
      }

      axios.patch(window.location.origin + "/api/sistem/clock/password",{id:1,password_clock:this.passwordExcel})
        .then(response => {
          if(response.data.status == "success")
          {             
            this.$notify({
              group: 'notifikasi',                  
              title: 'Update Password Sukses',
              type: 'success',
              text: 'Password Clock '+this.passwordExcel+' Berhasil Di Update'
            });
          }     
        }).catch(error => {
          this.$notify({
              group: 'notifikasi',                  
              title: 'Update Data Gagal',
              type: 'error',
              text: error.response.data.message
          }); 
        })
    },
    gantiPasswordPiket(){
      if(this.password == ''){
        this.$notify({
              group: 'notifikasi',                  
              title: 'Update Data Gagal',
              type: 'error',
              text: "Periksa Kembali Input Data"
        }); 
        return
      }

      axios.patch(window.location.origin + "/api/sistem/piket/password",{id:1,password_piket:this.password})
        .then(response => {
          if(response.data.status == "success")
          {            
            this.$notify({
              group: 'notifikasi',                  
              title: 'Update Password Sukses',
              type: 'success',
              text: 'Password Piket '+this.password+' Berhasil Di Update'
            });
          }     
        }).catch(error => {
          this.$notify({
              group: 'notifikasi',                  
              title: 'Update Data Gagal',
              type: 'error',
              text: error.response.data.message
          }); 
        })
    }
    //eend
  }  
  //----
}
</script>