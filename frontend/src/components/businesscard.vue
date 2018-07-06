<template>
  <div class="businesscard">
    <h1>{{ title }}</h1>
    <h3>Поиск визиток по ФИО</h3>
    <input type="text" placeholder="Поиск" v-model="search" />
    <table>
      <tr>
        <th>ФИО</th>
        <th>Должность</th>
        <th>Телефон</th>
        <th>E-mail</th>
        <th>Место работы</th>
        <th>Адрес</th>
      </tr>
      <tr class="data" v-for="vcard in filteredVcards" v-bind:key="vcard.phone">
        <td>{{vcard.name}}</td>
        <td>{{vcard.position}}</td>
        <td>{{vcard.phone}}</td>
        <td>{{vcard.email}}</td>
        <td>{{vcard.workplace}}</td>
        <td>{{vcard.address}}</td>
        <td><button
        v-on:click="edit_vcard(vcard)">Редактировать визитку</button></td>
        <td><button
        v-on:click="remove_vcard(vcard)">Удалить визитку</button></td>
      </tr>
    </table>
    <h3 v-if="edit_index == -1">Новая визитка</h3>
    <form>
      <p>ФИО: <input type="text" v-model="new_vcard.name"></p>
      <p>Должность: <input type="text" v-model="new_vcard.position"></p>
      <p>Телефон: <input type="text" v-model="new_vcard.phone"></p>
      <p>E-mail: <input type="text" v-model="new_vcard.email"></p>
      <p>Место работы: <input type="text" v-model="new_vcard.workplace"></p>
      <p>Адрес: <input type="text" v-model="new_vcard.address"></p>
      <button v-on:click="add_new_vcard" v-if="edit_index == -1">Добавить визитку</button>
      <button v-on:click="make_new_vcard" v-if="edit_index > -1">Закончить редактирование</button>
    </form>
  </div>
</template>

<script>
export default {
  name: 'businesscard',
  props: {
    title: String
  },
  data: function () {
    return {
      vcard_list: [
        {
          'name': 'Иванов Иван Иванович',
          'position': 'Директор',
          'phone': '+7-999-999-99-98',
          'email': 'ivanov@mail.ru',
          'workplace': 'ООО "Эврика"',
          'address': 'г.Красноярск, ул.Ленина, д.1'
        },
        {
          'name': 'Петров Пётр Петрович',
          'position': 'Водитель',
          'phone': '+7-999-999-99-99',
          'email': 'petrov@mail.ru',
          'workplace': 'ООО "Эврика"',
          'address': 'г.Красноярск, ул.Ленина, д.1'
        }
      ],
      new_vcard:
        {
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
  methods: {
    add_new_vcard: function () {
      this.vcard_list.push(this.new_vcard)
    },
    remove_vcard: function (item) {
      this.vcard_list.splice(this.vcard_list.indexOf(item), 1)
    },
    edit_vcard: function (item) {
      this.edit_index = this.vcard_list.indexOf(item)
      this.new_vcard = this.vcard_list[this.edit_index]
    },
    make_new_vcard: function () {
      this.edit_index = -1
      this.new_vcard = {
        'name': '',
        'position': '',
        'phone': '',
        'email': '',
        'workplace': '',
        'address': ''
      }
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
<style scoped>
</style>
