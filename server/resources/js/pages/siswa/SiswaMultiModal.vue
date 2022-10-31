<template>
<modal name="siswa-multi" draggable=".card-header" :width="400" :height="199.9" @before-open="beforeOpen" @before-close="beforeClose">
  <div class="card">
    <div class="card-header card-header-modal">
     <h5 style="position: absolute; margin-top: 15px;">Multi Edit</h5>
      <div class="card-options">
       <button class="close" @click="$modal.hide('siswa-multi')">
          <i class="fe fe-x text-muted"></i>
       </button>
      </div>
    </div>
    <div class="card-body card-body-modal">
    <div class="form-row">
      <div class="col-md-6">
        <label for="kelas">Kelas</label>
        <v-select :options="kelasData" label="nama" :reduce="nama => nama.id" v-model="kelas" placeholder="Semua Kelas" style="position: fixed;width: 13.5%;"></v-select>        
      </div>
      <div class="col-md-6">
        <label for="kelas">Status</label>
        <select class="form-control" v-model="status">
          <option value="a">Aktif</option>
          <option value="l">Lulus</option>          
          <option value="p">Pindah Sekolah</option>
          <option value="d">Drop Out</option>
          <option value="m">Meninggal</option>
        </select> 
      </div> 
    </div>
      <div class="float-right" style="margin-top:20px;">
        <button class="btn btn-primary" @click="updateSiswa()">Update Data</button>
      </div> 
    </div>
  </div>   
  </modal>
</template>
<script type="application/javascript">
export default {
  name: 'SiswaMultiModal',
  data () {
    return {                
      status:'a',
      kelas:'',
      kelasData:[]
    }
  },
  methods: {  
    fetchKelas(){
      axios.get(window.location.origin+"/api/allkelas")
      .then(response => {           
          this.kelasData = response.data.data          
        }).catch(error => this.kelasData = [])
    },    
    updateSiswa:function(event){
      
        axios.patch(window.location.origin+'/api/multisiswa', {
          selected:this.$parent.selected ,status: this.status, kelas_id: Number(this.kelas)
        })
        .then(response => {
          this.$parent.fetchData();
          if(response.data.status == "success")
          { 
            //console.log('berhasil diupdate');
            this.$modal.hide('siswa-multi')
            if(this.status === "a"){
              this.$notify({
                group: 'notifikasi',                  
                title: 'Update Siswa Sukses',
                type: 'success',
                text: 'Data Siswa Berhasil Di Update'
              });
            }else{
              this.$notify({
                group: 'notifikasi',                  
                title: 'Update Siswa Sukses',
                type: 'warn',
                text: 'Data Siswa Berhasil Di Arsipkan'
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
      this.fetchKelas()
      this.kelas = this.$parent.kelas;       
    },
    beforeClose (event) {
      //console.log("close")            
    }
   //end method
  }
}
</script>