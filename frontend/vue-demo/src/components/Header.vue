<template>
  <div class="header">
    <b-navbar type="light" variant="">
      <h2 class="header-title">タスク管理アプリ</h2>
      <b-collapse is-nav>
        <b-navbar-nav class="ml-auto">
          <b-nav-item><span>ようこそ　{{ getCurrentUserEmail }}</span></b-nav-item>
          <b-nav-item-dropdown text="メニュー" right>
            <b-dropdown-item @click="signOut">サインアウト</b-dropdown-item>
          </b-nav-item-dropdown>
        </b-navbar-nav>
      </b-collapse>
    </b-navbar>
  </div>
</template>

<script>
import firebase from 'firebase'
export default {
  data: {
    currentUserEmail: "",
  },
  methods: {
    signOut() {
      firebase.auth().signOut()
      localStorage.removeItem('jwt')
      this.$router.push('/signin')
    },
    getCurrentUserEmail() {
      if (firebase.auth().currentUse !== null) {
        this.getCurrentUserEmail = firebase.auth().currentUser.email
      }
    },
  },
  created() {
    this.getCurrentUserEmail()
  },
}
</script>

<style lang="scss" scoped>
.header {
  color: #333;
  background: #e1f9f6;
  text-align: center;
  padding: 5px 0;
}
.header-title {
  text-align: center;
  font-size: 20px;
  font-weight: bold;
  position: absolute;
  width: 100%;
  left: 0;
  margin: 0 auto;
  z-index: auto;
}
</style>