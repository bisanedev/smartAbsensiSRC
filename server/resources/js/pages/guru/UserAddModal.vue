<template>
<modal name="user-add" draggable=".card-header" :width="400" :height="600.3" @before-close="beforeClose">
  <div class="card">
    <div class="card-header card-header-modal">
     <h5 style="position: absolute; margin-top: 15px;">Input Guru</h5>
      <div class="card-options">
       <button class="close" @click="$modal.hide('user-add')">
          <i class="fe fe-x text-muted"></i>
       </button>
      </div>
    </div>
    <div class="card-body card-body-modal">
    <div class="form-row">
      <div class="col-md-12">
        <label for="name">Nama</label>
        <input type="text" class="form-control" placeholder="Nama Lengkap" v-model="nama">
      </div>
      <div class="col-md-6">
        <label for="mapel">Mapel</label>
        <input type="text" class="form-control" placeholder="Mata Pelajaran" v-model="mapel">
      </div>
      <div class="col-md-6">
        <label for="warna">Warna Label</label>
        <colorpicker :color="defaultColor" v-model="defaultColor" />        
      </div>
      <div class="col-md-12">
        <label for="jenis">Jenis Kelamin</label>
        <select class="form-control" v-model="jenis">
          <option>Laki-Laki</option>
          <option>Perempuan</option>          
        </select>                
      </div>
      <div class="col-md-12">
        <label for="username">Username</label>
        <input type="text" @keydown.space.prevent class="form-control" placeholder="Username" v-model="username">
      </div>
      <div class="col-md-12">
        <label for="Email">Role / Akses</label>
        <select class="form-control" v-model="role">
          <option>Guru</option>
          <option>Admin</option>          
        </select>                
      </div>
      <div class="col-md-12">
        <label for="password">Passward</label>
        <input type="password" class="form-control" placeholder="Password" v-model="password">
      </div>
      <div class="col-md-12">
        <label for="password">Re-passward</label>
        <input type="password" class="form-control" placeholder="Re-password" v-model="repassword">
      </div>   
    </div>
      <div class="float-right" style="margin-top:20px;">
        <button class="btn btn-primary" :disabled="disable" @click="adduser()">Input Data</button>
      </div> 
    </div>
  </div>   
  </modal>
</template>
<script type="application/javascript">
import colorpicker from '../../components/colorpicker.vue'
export default {
  name: 'UserAddModal',
  components : {
    colorpicker
  },
  data () {
    return {  
      disable: false,   
      username: '',
      jenis:'',
      nama:'',
      mapel:'',
      defaultColor: '',
      role:'Guru',      
      password:'',
      repassword:''
    }
  },
  methods: {      
    adduser:function(event){
        this.disable = true;
      if(this.password == this.repassword){
        axios.post(window.location.origin+'/api/users', {
        username: this.username, status: "a", name: this.nama, jenis: this.jenis, mapel: this.mapel, color: this.defaultColor, role: this.role, password: this.password
        })
        .then(response => {
          this.$parent.fetchData();
          if(response.data.status == "success")
          { 
            //console.log('berhasil diupdate');
            this.$modal.hide('user-add')
            this.$notify({
              group: 'notifikasi',                  
              title: 'Input Guru Sukses',
              type: 'success',
              text: 'Data Guru/Admin '+this.username+' Berhasil Di Tambahkan'
            });
          }
        })
        .catch(error => {
            this.disable = false;
            //console.log('gagal');
            this.$notify({
              group: 'notifikasi',                  
              title: 'Input Data Gagal',
              type: 'error',
              text: error.response.data.message
            });      
          
        })

      }else{
          this.$notify({
            group: 'notifikasi',                  
            title: 'Input Data Gagal',
            type: 'alert',
            text: "Please Confirm Password "
          }); 
      }      
    },
    beforeClose (event) {
      this.disable = false;
      //console.log("close")
      this.nama = this.aktif = this.defaultColor = this.jenis = this.username = this.password = this.repassword = this.mapel = '';      
    }
   //end method
  }
}
</script>