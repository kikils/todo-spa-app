<template>
  <div class="card" style="width: 18rem">
    <div v-if="validationError">タイトルを入力してください</div>
    <div v-if="apiError">サーバーでエラーが発生しました</div>
    <b-list-group class="list-group list-group-flush">
      <b-list-group-item class="list-group-item">
        <label>タイトル</label>
        <input
          type="text"
          name="title"
          class="form-control"
          v-model="inputTitle"
        />
      </b-list-group-item>
      <b-list-group-item class="list-group-item">
        <label>締め切り日付</label>
        <input
          type="text"
          name="period"
          class="form-control"
          v-model="inputPeriod"
        />
      </b-list-group-item>
      <b-list-group-item class="list-group-item">
        <label>詳細</label>
        <textarea
          name="detail"
          class="form-control"
          v-model="inputDetail"
        ></textarea>
      </b-list-group-item>
    </b-list-group>
    <button class="btn btn-success" v-on:click="callCreateTask">作成</button>
  </div>
</template>

<script>
import api from "@/api";
export default {
  data() {
    return {
      inputTitle: "",
      inputPeriod: "",
      inputDetail: "",
      validationError: false,
      apiError: false,
    };
  },
  methods: {
    callCreateTask() {
      if (this.inputTitle === "") {
        if (this.inputTitle === "") {
          this.validationError = true;
        }
        return;
      } else {
        this.validationError = false;
      }
      const todo = {
        title: this.inputTitle,
        note: this.inputDetail,
        due_date: this.inputPeriod,
        is_completed: false,
      };
      this.$emit("createTask", todo);
      this.initInput();
    },
    initInput() {
      this.inputTitle = "";
      this.inputDetail = "";
      this.inputPeriod = "";
    },
  },
};
</script>