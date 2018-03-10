import api from '@/service/api'

const STORAGE_USERNAME = 'myapp.auth.username'
const STORAGE_TOKEN = 'myapp.auth.token'

const state = {
  username: localStorage.getItem(STORAGE_USERNAME),
  token: localStorage.getItem(STORAGE_TOKEN)
}

// getters
const getters = {
  loggedIn (state) {
    return !!state.token
  }
}

// actions
const actions = {
  authenticate ({ commit, state }, { username, password }) {
    return new Promise((resolve, reject) => {
      api
        .token(username, password)
        .then(response => {
          commit('username', response.data.username)
          commit('token', response.data.token)
          resolve()
        }, reject)
    })
  },
  logout ({ commit, state }) {
    return new Promise((resolve, reject) => {
      api
        .logout(state.username, state.token)
        .then(response => {
          commit('username', '')
          commit('token', '')
          resolve()
        }, reject)
    })
  }
}

// mutations
const mutations = {
  username (state, username) {
    state.username = username
    localStorage.setItem(STORAGE_USERNAME, username)
  },

  token (state, token) {
    state.token = token
    localStorage.setItem(STORAGE_TOKEN, token)
  }
}

export default {
  state,
  getters,
  actions,
  mutations
}
