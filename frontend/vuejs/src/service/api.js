import axios from 'axios'

const ENDPOINT = 'http://127.0.0.1:3000'

export default {
  token (username, password) {
    const data = {
      username,
      password
    }

    const url = `${ENDPOINT}/token`

    return axios.post(url, data)
  },

  logout (username, token) {
    const data = {
      username,
      token
    }

    const url = `${ENDPOINT}/token`

    return axios.delete(url, data)
  },

  users (accessToken, { keyword, since, limit }) {
    keyword = keyword || ''
    since = since || ''
    limit = limit || 15

    const url = `${ENDPOINT}/users?keyword=${keyword}&since=${since}`

    return axios.get(url, {
      headers: {'X-Token': accessToken}
    })
  }
}
