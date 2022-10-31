<template>
<modal name="user-edit" draggable=".card-header" :width="400" :height="466.9" @before-open="beforeOpen" @before-close="beforeClose">
  <div class="card">
    <div class="card-header card-header-modal">
     <h5 style="position: absolute; margin-top: 15px;">Edit Data : <b>{{username}}</b></h5>
      <div class="card-options">
       <button class="close" @click="$modal.hide('user-edit')">
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
      <div class="col-md-6">
        <label for="username">Username</label>
        <input type="text" @keydown.space.prevent class="form-control" placeholder="Username" v-model="username">
      </div>
      <div class="col-md-6" v-if="id != '1'">
        <label for="kelas">Status</label>
        <select class="form-control" v-model="status">
          <option value="a">Aktif</option>
          <option value="p">Pensiun</option>
          <option value="b">Berhenti</option>          
          <option value="m">Meninggal</option>
        </select> 
      </div>
      <div class="col-md-12">
        <label for="Email">Role / Akses</label>
        <select class="form-control" v-model="role" v-if="role == 'superuser' ">
          <option>superuser</option>
        </select>
        <select class="form-control" v-model="role" v-else>
          <option>guru</option>
          <option>admin</option>          
        </select>                
      </div>
    </div>
      <div class="float-right" style="margin-top:20px;">
        <button class="btn btn-primary" @click="updateUser()">Update Data</button>
      </div> 
    </div>
  </div>   
  </modal>
</template>
<script type="application/javascript">
import colorpicker from '../../components/colorpicker.vue'
export default {
  name: 'UserEditModal',
  components : {
    colorpicker
  },
  data () {
    return {
      id:'',
      status:'',
      username: '',
      nama:'',
      status:'',      
      jenis:'',
      mapel:'',
      defaultColor:'',
      role:''
    }
  },
  methods: {      
    updateUser:function(event){
      
        axios.patch(window.location.origin+'/api/users/'+this.id, {
        username: this.username, status: this.status, name: this.nama, jenis: this.jenis, mapel: this.mapel, color:this.defaultColor, role: this.role
        })
        .then(response => {
          this.$parent.fetchData();
          if(response.data.status == "success")
          { 
            //console.log('berhasil diupdate');
            this.$modal.hide('user-edit')
            if(this.status === "a"){
                this.$notify({
                  group: 'notifikasi',                  
                  title: 'Update Guru/Admin Sukses',
                  type: 'success',
                  text: 'Data Guru/Admin '+this.username+' Berhasil Di Update'
                });
            }else{
              this.$notify({
                group: 'notifikasi',                  
                title: 'Update Guru/Admin Sukses',
                type: 'warn',
                text: 'Data Guru/Admin '+this.username+' Berhasil Di Arsipkan'
              });
            }
          }
        })
        .catch(error => {
        
            //console.log('gagal');
            this.$notify({
              group: 'notifikasi',                  
              title: 'Update Gagal',
              type: 'error',
              text: error.response.data.message
            });      
          
        })    
    },
    beforeOpen (event) {   
      this.id = event.params.wow.id;      
      this.nama = event.params.wow.nama;
      this.status = event.params.wow.status;
      this.jenis = event.params.wow.jenis;
      this.username = event.params.wow.username;
      this.mapel = event.params.wow.mapel;
      this.role = event.params.wow.role;
      this.defaultColor = event.params.wow.color;
    },
    beforeClose (event) {
      //console.log("close")         
    }
   //end method
  }
}
</script>