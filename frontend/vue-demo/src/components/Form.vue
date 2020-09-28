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
            placeholder="タイトル"
          />
        </b-list-group-item>
        <b-list-group-item class="list-group-item">
          <label>締め切り日付</label>
          <Datepicker v-model="inputPeriod" format="yyyy/MM/dd" :language="ja" input-class="form-control" :typeable=true placeholder="2020/09/27"></Datepicker>
        </b-list-group-item>
        <b-list-group-item class="list-group-item">
          <label>詳細</label>
          <textarea
            name="detail"
            class="form-control"
            v-model="inputDetail"
            placeholder="詳細"
          ></textarea>
        </b-list-group-item>
        <b-list-group-item class="list-group-item m-0 p-0">
          <div class="text-center p-2">
            <div v-if="editTodoId === 0">
              <button class="btn btn-success btn-block" v-on:click="callCreateTask">
                <span>作成</span>
              </button>
            </div>
            <div v-if="editTodoId !== 0">
              <button class="btn btn-danger btn-block" v-on:click="initInput">
                <span>編集をやめる</span>
              </button>
              <button class="btn btn-success btn-block" v-on:click="callCreateTask">
                <span>更新</span>
              </button>
            </div>
          </div>
        </b-list-group-item>
      </b-list-group>
    </div>
  </div>
</template>

<script>
import Datepicker from 'vuejs-datepicker'
import {ja} from 'vuejs-datepicker/dist/locale'

import api from "@/api";
export default {
  components: {
    Datepicker
  },
  props: ['editTodoId'],
  data() {
    return {
      inputTitle: "",
      inputPeriod: "",
      inputDetail: "",
      validationError: "",
      apiError: false,
      ja:ja,
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
      if (this.todo_id == 0) {
        const todo = {
          title: this.inputTitle,
          note: this.inputDetail,
          due_date: this.inputPeriod,
          is_completed: false,
        };
        this.$emit("createTask", todo);
      } else {
        const todo = {
          id: this.editTodoId,
          title: this.inputTitle,
          note: this.inputDetail,
          due_date: this.inputPeriod,
          is_completed: false,
        };
        this.$emit("updateTodo", todo);
        this.editTodoId = 0
        this.$emit("updateEditTodoId", 0)
      }
      this.initInput();
    },
    setTodo(todo) {
      this.editTodoId = todo.id
      this.$emit("updateEditTodoId", todo.id)
      this.inputTitle = todo.title;
      this.inputDetail = todo.note;
      this.inputPeriod = todo.due_date;
    },
    initInput() {
      this.editTodoId = 0
      this.$emit("updateEditTodoId", 0)
      this.inputTitle = "";
      this.inputDetail = "";
      this.inputPeriod = "";
    },
  },
};
</script>