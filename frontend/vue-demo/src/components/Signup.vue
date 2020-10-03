<template>
  <div class="text-center">
    <div class="signup">
      <h2 class="h3 my-3 font-weight-normal">サインアップ</h2>
      <input
        type="email"
        placeholder="email"
        class="form-control"
        v-model="email"
        required
      />
      <input
        type="password"
        placeholder="Password"
        class="form-control"
        v-model="password"
      />
      <button
        class="btn btn-lg btn-primary btn-block"
        type="submit"
        @click="signUp"
      >
        登録する
      </button>
      <p>
        アカウントを持っている場合は
        <router-link to="/signin">こちら</router-link>
      </p>
    </div>
  </div>
</template>

<script>
import firebase from "firebase";
export default {
  name: "Signup",
  data() {
    return {
      email: "",
      password: "",
    };
  },
  methods: {
    signUp: function() {
      firebase
        .auth()
        .createUserWithEmailAndPassword(this.email, this.password)
        .then((res) => {
          this.$router.push("/signin");
        })
        .catch((error) => {
          alert(error.message);
        });
    },
  },
};
</script>

<style lang="scss" scoped>
.signup {
  width: 100%;
  max-width: 330px;
  padding: 15px;
  margin: 0 auto;
}
</style>
