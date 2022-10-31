<template>
<modal name="semester-edit" draggable=".card-header" :width="400" :height="400.9" @before-open="beforeOpen" @before-close="beforeClose">
  <div class="card">
    <div class="card-header card-header-modal">
     <h5 style="position: absolute; margin-top: 15px;">Edit Semester</h5>
      <div class="card-options">
       <button class="close" @click="$modal.hide('semester-edit')">
          <i class="fe fe-x text-muted"></i>
       </button>
      </div>
    </div>
    <div class="card-body card-body-modal">
    <div class="form-row">  
      <div class="col-md-12">
        <label for="nama">Tahun Ajaran</label>
         <input type="text" class="form-control" placeholder="Tahun Ajaran" v-model="semester">
      </div>
      <div class="col-md-12">
        <label for="nama">Pilih Semester</label>
        <select class="form-control" v-model="semesterdetail">
          <option>Ganjil</option>
          <option>Genap</option>          
        </select>           
      </div>
      <div class="col-md-12">
        <label for="mapel">Tanggal Awal Semester</label>
        <date-picker v-model="waktustart" value-type="YYYY-MM-DD" format="DD/MM/YYYY" type="date" placeholder="Pilih Tanggal">                          
        </date-picker>     
      </div>
      <div class="col-md-12">
        <label for="mapel">Tanggal Akhir Semester</label>
        <date-picker v-model="waktuend" value-type="YYYY-MM-DD" format="DD/MM/YYYY" type="date" placeholder="Pilih Tanggal">                          
        </date-picker>       
      </div>
    </div>
      <div class="float-right" style="margin-top:20px;">
        <button class="btn btn-primary" :disabled="disable" @click="addsemester()">Input Data</button>
      </div> 
    </div>
  </div>   
  </modal>
</template>
<script type="application/javascript">
import moment from 'moment';
export default {
  name: 'SemesterEditModal',
  data () {
    return { 
      disable: false,         
      semester:'',
      semesterdetail:'',
      waktustart:'',
      waktuend:'',
    }
  },
  methods: {      
    addsemester:function(event){
        if (this.semesterdetail == '' && this.semester == '' && this.waktustart == '' && this.waktuend == ''){
           this.$notify({
              group: 'notifikasi',                  
              title: 'Update Data Gagal',
              type: 'error',
              text: 'Periksa Kembali Input Data'
            });
          return
        }
        this.disable = true;
        axios.patch(window.location.origin+'/api/semester/'+this.id, {
            nama: this.semester,semester:this.semesterdetail, waktustart: this.waktustart, waktuend: this.waktuend
        })
        .then(response => {
          this.$parent.fetchData();
          if(response.data.status == "success")
          {             
            //console.log('berhasil diupdate');
            this.$modal.hide('semester-edit')
            this.$notify({
              group: 'notifikasi',                  
              title: 'Update Semester Sukses',
              type: 'success',
              text: 'Data Semester '+this.semester+' Berhasil Di Update'
            });
          }
        })
        .catch(error => {        
            //console.log('gagal');
            this.$notify({
              group: 'notifikasi',                  
              title: 'Update Data Gagal',
              type: 'error',
              text: error.response.data.message
            });      
          
        })     
    },
     beforeOpen (event) {
      this.id = event.params.wow.id;
      this.semester = event.params.wow.nama;
      this.semesterdetail = event.params.wow.semester;
      this.waktustart = moment(event.params.wow.waktustart).format('YYYY-MM-DD'); 
      this.waktuend = moment(event.params.wow.waktuend).format('YYYY-MM-DD');      
    },
    beforeClose (event) {
      this.disable = false;
      //console.log("close")      
      //this.semester = this.waktuend = this.waktustart = this.semesterdetail = '';     
    }
   //end method
  }
}
</script>