<script src="./style.css" type="text/css"></script>
<template>
  <div class="businesscard">
    <b-jumbotron text-variant="black" border-variant="dark"
    style="background-color: #fd8aa5; color: #38000d; text-align: center;
    font-family: 'Segoe Print'">
      <template slot="header">
        {{ title}}
      </template>
    </b-jumbotron>
    <h3 v-if="error">Ошибка: {{ error }}</h3>
    <div >
     <h3 style="text-align:center" class="h3Search">Поиск визиток по ФИО</h3>
     </div>
     <div class="form-control1">
    <input type="text" placeholder="Чтобы найти визитку, введите ФИО" v-model="search" class="form-control"
    style="margin: auto; margin-top: 10px"/></div>
    <div>
      <div v-for="vcard in filteredVcards" v-bind:key="vcard.id">
        <div class="vcard">
          <div class="place">{{vcard.workplace}}</div><hr class="hr1">
          <h2>{{vcard.name}}</h2>
          {{vcard.position}}<br><hr>
          <div class="left">{{vcard.address}}</div>
          <div class="right">телефон: {{vcard.phone}}<br>
          e-mail: {{vcard.email}}</div>
        </div>
        <br>
        <div class="buttons">
          <div class="left1">
            <button v-on:click="edit_vcard(vcard)" type="button" class="btn btn-primary">Редактировать визитку</button>
          </div>
          <div class="right1">
            <button v-on:click="remove_vcard(vcard)" type="button" class="btn btn-primary">Удалить визитку</button>
          </div>
        </div>
        <br>
      </div>
    </div>
    <br>
    <h3 v-if="filteredVcards.length === 0 && vcard_list.length != 0" class="searchWarning">
      Нет визиток с введёнными данными!</h3>
      <br>
    <div>
    <b-card class="newVcard">
    <b-form>
    <h3 v-if="edit_index == -1" class="h3Form">Новая визитка</h3>
      <b-form-group label="ФИО:" class="formLabel">
        <b-form-input v-model="new_vcard.name" required class="form-control" placeholder="Введите ФИО">
        </b-form-input>
      </b-form-group>
      <b-form-group label="Должность:" class="formLabel">
        <b-form-input v-model="new_vcard.position" required class="form-control" placeholder="Введите должность">
        </b-form-input>
      </b-form-group>
      <b-form-group label="Телефон:" class="formLabel">
        <b-form-input v-model="new_vcard.phone" required class="form-control" placeholder="Введите номер телефона">
        </b-form-input>
      </b-form-group>
      <b-form-group label="Email:" class="formLabel">
        <b-form-input v-model="new_vcard.email" type="email" required class="form-control" placeholder="Введите Email">
        </b-form-input>
      </b-form-group>
      <b-form-group label="Место работы:" class="formLabel">
        <b-form-input v-model="new_vcard.workplace" required class="form-control"
        placeholder="Введите название организации">
        </b-form-input>
      </b-form-group>
      <b-form-group label="Адрес:" class="formLabel">
        <b-form-input v-model="new_vcard.address" required class="form-control" placeholder="Введите адрес">
        </b-form-input>
      </b-form-group>
      <b-button v-if="edit_index == -1" v-on:click="add_new_vcard" type="button"
      style="margin-left:210px !important"
      class="btn btn-primary">Добавить визитку</b-button>
      <b-button v-if="edit_index > -1" v-on:click="end_of_edition" type="button"
      style="margin-left:210px !important"
      class="btn btn-primary">Закончить редактирование</b-button>
    </b-form>
    </b-card>
    </div>
    <br><br><br><br>
  </div>
</template>

<script>
const axios = require('axios')

export default {
  name: 'businesscard',
  props: {
    title: String
  },
  data: function () {
    return {
      vcard_list: [],
      error: '',
      new_vcard:
        {
          'id': 0,
          'name': '',
          'position': '',
          'phone': '',
          'email': '',
          'workplace': '',
          'address': ''
        },
      edit_index: -1,
      search: ''
    }
  },
  mounted: function () {
    this.get_vcards()
  },
  methods: {
    get_vcards: function () {
      this.error = ''
      const url = '/api/businesscard/vcards'
      axios.get(url).then(response => {
        this.vcard_list = response.data
      }).catch(response => {
        this.error = response.response.data
      })
    },
    add_new_vcard: function () {
      const url = '/api/businesscard/vcards'
      axios.post(url, this.new_vcard).then(response => {
        this.vcard_list.push(this.new_vcard)
      }).catch(response => {
        this.error = response.response.data
      })
    },
    remove_vcard: function (item) {
      const url = '/api/businesscard/vcards/' + item.id
      axios.delete(url).then(response => {
        this.vcard_list.splice(this.vcard_list.indexOf(item), 1)
      }).catch(response => {
        this.error = response.response.data
      })
    },
    edit_vcard: function (item) {
      this.edit_index = this.vcard_list.indexOf(item)
      this.new_vcard = this.vcard_list[this.edit_index]
    },
    end_of_edition: function () {
      const url = '/api/businesscard/vcards/' + this.new_vcard.id
      axios.put(url, this.new_vcard).then(response => {
        this.edit_index = -1
        this.new_vcard = {
          'id': 0,
          'name': '',
          'position': '',
          'phone': '',
          'email': '',
          'workplace': '',
          'address': ''
        }
      }).catch(response => {
        this.error = response.response.data
      })
    }
  },
  computed: {
    filteredVcards: function () {
      return this.vcard_list.filter(function (item) {
        return item.name.indexOf(this.search) !== -1
      }.bind(this))
    }
  }
}
</script>
<style>
</style>
