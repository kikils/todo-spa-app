<template>
  <div id="form">
    <div v-if="validationError" class="text-left text-danger">{{ validationError }}</div>
    <div class="card" style="width: 18rem">
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
        <b-list-group-item class="list-group-item m-0 p-0">
          <div class="text-center p-2">
            <button class="btn btn-success btn-block" v-on:click="callCreateTask">
              作成
            </button>
          </div>
        </b-list-group-item>
      </b-list-group>
    </div>
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
      validationError: "",
      apiError: false,
    };
  },
  methods: {
    callCreateTask() {
      if (this.inputTitle === "") {
          this.validationError = "タイトルを入力してください"
          return
      }
      if (this.inputPeriod === "") {
          this.validationError = "締め切り日付を入力してください"
          return
      }
      this.validationError = ""
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