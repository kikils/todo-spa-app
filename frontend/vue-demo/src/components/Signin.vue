<template>
  <div class="text-center">
    <div class="signin">
      <h2 class="h3 my-3 font-weight-normal">サインイン</h2>
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
        @click="signIn"
      >
        Sign in
      </button>
      <p>
        アカウントを持っていない場合は
        <router-link to="/signup">こちら</router-link>
      </p>
    </div>
  </div>
</template>

<script>
import firebase from "firebase";
export default {
  name: "Signin",
  data: function () {
    return {
      email: "",
      password: "",
    };
  },
  methods: {
    signIn: function () {
      firebase
        .auth()
        .signInWithEmailAndPassword(this.email, this.password)
        .then(
          (res) => {
            res.user.getIdToken().then((idToken) => {
              localStorage.setItem("jwt", idToken.toString());
            });
            this.$router.push("/");
          },
          (err) => {
            alert(err.message);
          }
        );
    },
  },
};
</script>

<style lang="scss" scoped>
.signin {
  width: 100%;
  max-width: 330px;
  padding: 15px;
  margin: 0 auto;
}
</style>