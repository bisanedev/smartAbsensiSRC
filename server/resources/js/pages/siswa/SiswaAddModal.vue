<template>
<modal name="siswa-add" draggable=".card-header" :width="400" :height="399.6" @before-open="beforeOpen" @before-close="beforeClose">
  <div class="card">
    <div class="card-header card-header-modal">
     <h5 style="position: absolute; margin-top: 15px;">Input Siswa</h5>
      <div class="card-options">
       <button class="close" @click="$modal.hide('siswa-add')">
          <i class="fe fe-x text-muted"></i>
       </button>
      </div>
    </div>
    <div class="card-body card-body-modal">
    <div class="form-row">
      <div class="col-md-12">
        <label for="kelas">Kelas</label>
        <v-select :options="kelasData" label="nama" :reduce="nama => nama.id" v-model="kelas" placeholder="Semua Kelas" @input="onChange($event)"></v-select>        
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
        <button class="btn btn-primary" :disabled="disable" @click="addsiswa()">Input Data</button>
      </div> 
    </div>
  </div>   
  </modal>
</template>
<script type="application/javascript">
export default {
  name: 'SiswaAddModal',
  data () {
    return { 
      disable: false,     
      urut:0,
      nama:'',
      jenis:'',
      kelas:'',
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
    onChange(event) {                            
      this.fetchJumlahSiswa()
    },
    fetchJumlahSiswa(){ 
      axios.get(window.location.origin+"/api/siswajumlah?kelas="+this.kelas)
      .then(response => {           
          this.urut = response.data.data + 1          
        }).catch(error => this.kelasData = [])
    }, 
    fetchKelas(){
      axios.get(window.location.origin+"/api/allkelas")
      .then(response => {           
          this.kelasData = response.data.data                  
        }).catch(error => this.kelasData = [])
    },      
    addsiswa:function(event){
        this.disable = true;
        axios.post(window.location.origin+'/api/siswa', {
        nama: this.nama, status:'a', jenis: this.jenis, kelas_id: Number(this.kelas) ,urut: Number(this.urut)
        })
        .then(response => {
          this.$parent.fetchData();
          if(response.data.status == "success")
          { 
            //console.log('berhasil diupdate');
            this.$modal.hide('siswa-add')
            this.$notify({
              group: 'notifikasi',                  
              title: 'Input Siswa Sukses',
              type: 'success',
              text: 'Data Siswa '+this.nama+' Berhasil Di Tambahkan'
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
    },
    beforeOpen (event) {
      this.fetchKelas()      
    },
    beforeClose (event) {
      this.disable = false;
      //console.log("close")
      this.urut = 0;
      this.nama = this.jenis = this.kelas = '';      
    }
   //end method
  }
}
</script>