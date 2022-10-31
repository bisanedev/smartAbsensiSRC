<template>
<modal name="siswa-edit" draggable=".card-header" :width="400" :height="399.6" @before-open="beforeOpen" @before-close="beforeClose">
  <div class="card">
    <div class="card-header card-header-modal">
     <h5 style="position: absolute; margin-top: 15px;">Edit : <b>{{nama}}</b></h5>
      <div class="card-options">
       <button class="close" @click="$modal.hide('siswa-edit')">
          <i class="fe fe-x text-muted"></i>
       </button>
      </div>
    </div>
    <div class="card-body card-body-modal">
    <div class="form-row">
      <div class="col-md-6">
        <label for="kelas">Kelas</label>
        <v-select :options="kelasData" label="nama" :reduce="nama => nama.id" v-model="kelas" placeholder="Semua Kelas"></v-select>        
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
      <div class="col-md-12">
        <label for="urut">No Absen</label>
        <div class="number-input">
          <button v-on:click="dec('urut', 1)" ></button>
          <input class="quantity" v-model.number="urut" type="number">
          <button v-on:click="inc('urut', 1)" class="plus"></button>          
        </div>     
      </div>
      <div class="col-md-12">
        <label for="name">Nama</label>
        <input type="text" class="form-control" placeholder="Nama Lengkap" v-model="nama">
      </div>
      <div class="col-md-12">
        <label for="jenis">Jenis Kelamin</label>
        <select class="form-control" v-model="jenis">
          <option>Laki-Laki</option>
          <option>Perempuan</option>          
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
  name: 'SiswaEditModal',
  data () {
    return {
      id:'',      
      nama:'',
      jenis:'',
      status:'',
      kelas:'',
      urut:0,
      kelasData:[]
    }
  },
  methods: {
    inc(property, amt){
      this[property] += amt
    },
    dec(property, amt){
      if (this[property] === 1) return
      this[property] -= amt
    },  
    fetchKelas(){
      axios.get(window.location.origin+"/api/allkelas")
      .then(response => {           
          this.kelasData = response.data.data          
        }).catch(error => this.kelasData = [])
    },    
    updateSiswa:function(event){
      
        axios.patch(window.location.origin+'/api/siswa/'+this.id, {
          nama: this.nama, status: this.status, jenis: this.jenis, kelas_id: Number(this.kelas), urut: Number(this.urut)
        })
        .then(response => {
          this.$parent.fetchData();
          if(response.data.status == "success")
          { 
            //console.log('berhasil diupdate');
            this.$modal.hide('siswa-edit')
            if(this.status === "a"){
              this.$notify({
                group: 'notifikasi',                  
                title: 'Update Siswa Sukses',
                type: 'success',
                text: 'Data Siswa '+this.nama+' Berhasil Di Update'
              });
            }else{
              this.$notify({
                group: 'notifikasi',                  
                title: 'Update Siswa Sukses',
                type: 'warn',
                text: 'Data Siswa '+this.nama+' Berhasil Di Arsipkan'
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
      this.id = event.params.wow.id;      
      this.nama = event.params.wow.nama;
      this.jenis = event.params.wow.jenis;                
      this.kelas = event.params.wow.kelas_id;
      this.urut = event.params.wow.urut;
      this.status = event.params.wow.status;
    },
    beforeClose (event) {
      //console.log("close")            
    }
   //end method
  }
}
</script>