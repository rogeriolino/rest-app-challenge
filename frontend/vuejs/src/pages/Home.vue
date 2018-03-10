<template>
  <div>
    <div class="container">
      <form>
        <div class="form-group">
          <label class="sr-only">Search</label>
          <input type="text" class="form-control" placeholder="Search" v-model="keyword" @keyup="changeKeyword">
        </div>
      </form>

      <button type="button" class="btn btn-outline-secondary" @click.prevent="prev" :disabled="loading">Previous</button>

      <button type="button" class="btn btn-outline-secondary" @click.prevent="next" :disabled="loading">Next</button>

      <span v-if="loading">loading...</span>

      <table class="table">
        <thead>
          <tr>
            <th>Username</th>
            <th>Name</th>
            <th></th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="user in users" :key="user.id">
            <td>{{user.username}}</td>
            <td>{{user.name}}</td>
            <td></td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script>
import api from '@/service/api'
import _ from 'lodash'

export default {
  name: 'Home',
  data () {
    return {
      keyword: '',
      users: [],
      loading: false,
      since: '',
      prevSince: []
    }
  },
  methods: {
    changeKeyword: _.debounce(function () {
      this.since = ''
      this.fetchUsers()
    }, 300),
    prev () {
      if (!this.loading && this.prevSince.length) {
        this.since = this.prevSince.pop()
        this.fetchUsers()
      }
    },
    next () {
      if (!this.loading && this.users.length) {
        this.prevSince.push(this.since)
        this.since = this.users[this.users.length - 1].username
        this.fetchUsers()
      }
    },
    fetchUsers () {
      if (this.loading) {
        return
      }
      this.loading = true
      api
        .users(this.$store.state.auth.token, {
          keyword: this.keyword,
          since: this.since
        })
        .then(response => {
          this.loading = false
          this.users = response.data
        }, () => {
          this.loading = false
        })
        .catch(() => {
          this.loading = false
        })
    }
  },
  beforeMount () {
    this.fetchUsers()
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
