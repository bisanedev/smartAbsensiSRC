<template>
    <div class="container">
        <div class="page-header">
            <h1 class="page-title">Arsip Guru</h1>
        </div>
        <div class="row row-cards row-deck">
            <div class="col-12">
                <div class="card">
                    <div class="card-header">
                        <div class="col-2" style="padding-left: 0px;">
                            <v-select :options="statusoption" label="label" :reduce="label => label.value" v-model="status" placeholder="Pilih Status" @input="onStatus()"></v-select>
                        </div> 
                        <div class="card-options">                         
                            <select class="form-control" v-model="perPage" @change="onChange($event)" style="width: 60px;margin-right:10px;">
                                <option>5</option>
                                <option>10</option>
                                <option>15</option>
                            </select>
                            <div class="input-group" style="width:260px;">
                            <input type="text" class="form-control" v-model="search" v-on:keyup="searchEnter" placeholder="Cari Nama">
                            <button class="btn bg-transparent" @click="refresh()" style="margin-left: -40px; z-index: 100;">
                                <i class="fe fe-x text-muted"></i>
                            </button>
                            <div class="input-group-append">
                                <button class="btn btn-primary" type="button" @click="cari()">
                                    <i class="fe fe-search"></i>
                                </button>
                            </div>
                            </div>
                        </div>
                    </div>
                    <div class="card-body" style="padding:0 !important;">
                        <div class="table-responsive">
                            <table class="table table-striped table-vcenter text-nowrap card-table">
                                <thead class="thead-dark">
                                    <tr>
                                        <th class="text-center w-1">#</th>                                     
                                        <th class="text-center w-1"><i class="icon-people"></i></th>
                                        <th @click="sortByColumn('nama')">
                                        Nama Guru
                                        <span v-if="'nama' === sortedColumn">
                                        <i v-if="order === 'asc' " class="fe fe-arrow-up"></i>
                                        <i v-else class="fe fe-arrow-down"></i>
                                        </span>
                                        </th> 
                                        <th class="text-center" @click="sortByColumn('jenis')">
                                        Jenis Kelamin
                                        <span v-if="'jenis' === sortedColumn">
                                        <i v-if="order === 'asc' " class="fe fe-arrow-up"></i>
                                        <i v-else class="fe fe-arrow-down"></i>
                                        </span>
                                        </th> 
                                        <th class="text-center" v-for="column in columns" :key="column" @click="sortByColumn(column)">
                                        {{ column | columnHead }}
                                        <span v-if="column === sortedColumn">
                                        <i v-if="order === 'asc' " class="fe fe-arrow-up"></i>
                                        <i v-else class="fe fe-arrow-down"></i>
                                        </span>
                                        </th>                                              
                                        <th class="text-center"></th>
                                    </tr>
                                </thead>
                                <tbody>
                                    <tr class="" v-if="tableData.length === 0">
                                        <td class="lead text-center" :colspan="columns.length + 5">No data found.</td>
                                    </tr>
                                    <tr v-else v-for="(data, key1) in tableData" :key="data.id">
                                        <td class="text-center">{{ serialNumber(key1) }}</td>
                                        <td class="text-center">
                                            <div v-if="data.foto" class="avatar d-block avfoto" v-bind:style="{ 'background-image':'url(./storage/guru/'+ data.foto +')' }">                                                
                                            </div>
                                            <div v-else-if="data.jenis == 'Perempuan'" class="avatar d-block avfoto" style="background-image: url(./assets/images/cewek.png)">
                                            </div>

                                            <div v-else class="avatar d-block avfoto" style="background-image: url(./assets/images/cowok.png)">
                                            </div>
                                        </td>
                                        <td>           
                                        {{ data.nama }}
                                        </td>
                                        <td class="text-center">           
                                        {{ data.jenis }}
                                        </td>
                                        <td v-if="data.mapel" class="text-center">
                                        {{ data.mapel }} <span class="mapel-label" v-bind:style="{ backgroundColor: data.color}"></span>
                                        </td>
                                        <td v-else class="text-center">
                                            <span class="badge badge-secondary">None</span>
                                        </td>
                                        <td class="text-center">         
                                        {{ data.username }}
                                        </td>
                                        <td class="text-center">
                                        {{ data.role }}
                                        </td>
                                        <td class="text-center">
                                            <div class="item-action dropdown">
                                                <a href="javascript:void(0)" data-toggle="dropdown" class="icon"><i class="fe fe-grid"></i></a>
                                                <div class="dropdown-menu dropdown-menu-right">
                                                    <button type="button" @click="uploadModal(data)" class="dropdown-item" ><i class="fe fe-image"></i> Foto </button>
                                                    <button type="button" @click="showEditModal(data)" class="dropdown-item"><i class="fe fe-edit"></i> Edit 
                                                    </button> 
                                                    <button type="button" @click="resetModal(data)" class="dropdown-item"><i class="fe fe-lock"></i> Change Password </button>
                                                    <div class="dropdown-divider"></div>
                                                    <button class="dropdown-item" @click="showDeleteDialog(data.id,data.username)"><i class="fe fe-trash"></i> Delete </a> </button>
                                                </div>
                                            </div>
                                        </td>
                                    </tr>
                                </tbody>
                            </table>
                        </div>
                    </div>
                    <div class="card-footer">
                      <span v-if="pagination.total_record" style="position: absolute;bottom: 20px;"> &nbsp; <i>Total Data : {{ pagination.total_record }} entries.</i></span>
                        <nav v-if="pagination && tableData.length > 0">
                        <ul class="pagination float-right">
                            <li class="page-item" :class="{'disabled' : currentPage === 1}">
                              <a class="page-link" href="#" @click.prevent="changePage(1)">
                                <i class="fe fe-chevrons-left"></i>
                              </a>
                            </li>
                            <li class="page-item" :class="{'disabled' : currentPage === 1}">
                              <a class="page-link" href="#" @click.prevent="changePage(currentPage - 1)"><i class="fe fe-chevron-left"></i></a>
                            </li>                            
                            <li v-for="page in pagesNumber" v-if="Math.abs(page - currentPage) < 3 || page == pagesNumber - 1 || page == 0" class="page-item" :class="{'active': page == pagination.page}">
                            <a href="javascript:void(0)" @click.prevent="changePage(page)" class="page-link">{{ page }}</a>
                            </li>                            
                            <li class="page-item" :class="{'disabled': currentPage === pagination.total_page}">
                            <a class="page-link" href="#" @click.prevent="changePage(currentPage + 1)"><i class="fe fe-chevron-right"></i></a>
                            </li>
                            <li class="page-item" :class="{'disabled': currentPage === pagination.total_page}">
                            <a class="page-link" href="#" @click.prevent="changePage(pagination.total_page)"><i class="fe fe-chevrons-right"></i></a>
                            </li>
                        </ul>
                        </nav>
                    </div>
                </div>
            </div>
        </div>
        <user-upload-modal/>
        <user-add-modal/>
        <user-edit-modal/>
        <reset-pass-modal/>
    </div>
</template>
<script type="application/javascript">

import UserUploadModal from './UserUploadModal.vue'
import UserAddModal from './UserAddModal.vue'
import UserEditModal from './UserEditModal.vue'
import ResetPassModal from './ResetPassModal.vue'

export default {
  components : {
    UserAddModal,
    UserUploadModal,
    UserEditModal,
    ResetPassModal
  },
  data() {
    return {      
      columns: ['mapel','username','role'],
      search:'',
      tableData: [],
      url: window.location.origin,
      pagination:[],
      statusoption:[{label: 'Pensiun', value: 'p'},{label: 'Berhenti', value: 'b'},{label: 'Meninggal', value: 'm'}],
      status: null,      
      offset: 4,
      currentPage: 1,
      perPage: '5',
      sortedColumn: 'nama',
      order: 'desc'
    }
  },
  created() {
    return this.fetchData()
  },
  computed: {   
    /**
     * Get the pages number array for displaying in the pagination.
     * */
    pagesNumber() {
      if (this.pagination.offset == undefined || this.pagination.offset == null) {
        return []
      }
      let from = this.pagination.page - this.offset
      if (from < 1) {
        from = 1
      }
      let to = from + (this.offset * 2)
      if (to >= this.pagination.total_page) {
        to = this.pagination.total_page
      }
      let pagesArray = []
      for (let page = from; page <= to; page++) {
        pagesArray.push(page)
      }
      return pagesArray
    },
    /**
     * Get the total data displayed in the current page.
     * */
    totalData() {
      return (this.pagination.offset - this.pagination.total_record) + 1
    }
  },
  methods: {
    fetchData() {
      let dataFetchUrl = `${this.url}/api/arsip/users?page=${this.currentPage}&search=${this.search}&order=${this.sortedColumn}&by=${this.order}&limit=${this.perPage}&status=${this.status}`
        axios.get(dataFetchUrl)
        .then(response => { 
          this.pagination = response.data
          this.tableData = response.data.records
        }).catch(error => this.tableData = [])
    },
    onStatus(){
      this.fetchData()
    },
    /**
     * Get the serial number.
     * @param key
     * */
    serialNumber(key) {
      return (this.currentPage - 1) * this.perPage + 1 + key
    },
    /**
     * Change the page.
     * @param pageNumber
     */
    changePage(pageNumber) {
      this.currentPage = pageNumber
      this.fetchData()
    },
    /** perpage ***/
    onChange(event) {
      this.currentPage = 1;
      this.fetchData();
    },
    /** enter search **/
    searchEnter:function(e){
      if (e.keyCode === 13) {
        this.currentPage = 1;
        this.fetchData() 
      }
    },
    /** cari **/
    cari(){
      this.currentPage = 1;
      this.fetchData()
    },
    /** refresh **/
    refresh(){
      this.search = "";
      this.fetchData();     
    },
    /** reset password **/
    resetModal(data){ 
      this.$modal.show('reset-pass',{wow:data});
    },
    /** Upload modal **/
    uploadModal(data){      
      this.$modal.show('user-upload',{wow:data});
    },
    /** User ADD **/
    showAddModal(){
      this.$modal.show('user-add');
    },
    /** edit modal **/
    showEditModal(data){      
      this.$modal.show('user-edit',{wow:data});
    },
    /** pass variable to modal **/ 
    showDeleteDialog (id,username) {
      this.$modal.show('dialog', {
        title: 'Konfirmasi',
        text: 'Apakah Anda benar-benar ingin menghapus? Username: '+username,
        buttons: [
          {
            title: 'Cancel',            
            handler: () => {
              this.$modal.hide('dialog')
            }
          },
          {
            title: 'Delete',
            default: true,
            handler: () => {                          
              axios.delete(this.url+'/api/users/'+id)
                .then(response => {                
                  if(response.data.status == "success")
                  { 
                    //console.log('berhasil dihapus');
                       
                    this.fetchData();
                    this.$modal.hide('dialog');
                    this.$notify({
                      group: 'notifikasi',                  
                      title: 'Delete Sukses',
                      type: 'success',
                      text: 'Data User '+username+' Berhasil Di Hapus'
                    });  
                  }
                })
                .catch(error => {
                    console.log(error);
                })
              //-------
            }
          }
        ]
      })
    },
    /**
     * Sort the data by column.
     * */
    sortByColumn(column) {
      if (column === this.sortedColumn) {
        this.order = (this.order === 'asc') ? 'desc' : 'asc'
      } else {
        this.sortedColumn = column
        this.order = 'asc'
      }
      this.fetchData()
    }
  },
  filters: {
    columnHead(value) {
      return value.split('_').join(' ').toUpperCase()
    }
  } 


//end script  
}
</script>