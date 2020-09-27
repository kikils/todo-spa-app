<template>
  <div id="app">
    <Header />
    <div class="container">
      <div class="row">
        <div class="col-md-12">
          <div v-if="apiError" class="mt-2 mb-2 p-2 alert alert-danger" role="alert">
            サーバーと通信ができません。しばらく時間をおいてからもう一度お試しください。
          </div>
          <div v-if="message" class="mt-2 mb-2 p-2 alert alert-success" role="alert">
            {{ message }}
          </div>
        </div>
        <div class="col-md-4">
          <Form @createTask="createTask" />
        </div>
        <div class="col-md-8">
          <List
            :todos="todos"
            :apiError="apiError"
            @deleteTodo="deleteTodo"
            @getTodo="getTodo"
            @completeTodo="completeTodo"
          />
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Vue from "vue";

import api from "@/api"

import Header from "./components/Header";
import Form from "./components/Form";
import List from "./components/List";
export default {
  name: "app",
  components: {
    Header,
    Form,
    List,
  },
  data() {
    return {
      todos: [],
      apiError: false,
      message: "",
    };
  },
  methods: {
    async getTodo() {
      let result;
      try {
        result = await api.get("/todo/all");
      } catch (err) {
        this.apiError = true;
        return;
      }
      this.todos = result.data;
      this.apiError = false;
      console.log(result);
    },
    async completeTodo(todo) {
      let result;
      try {
        result = await api.post("/todo/update", {
          id: todo.id,
          title: todo.title,
          note: todo.note,
          due_date: todo.due_date,
          is_completed: !todo.is_completed,
        });
      } catch (err) {
        this.apiError = true;
        return;
      }
      this.apiError = false;
      this.message = "正常に更新できました。"
      console.log(result);
      this.getTodo()
    },
    async deleteTodo(todoId) {
      let result;
      try {
        result = await api.post("/todo/delete", {
          id: todoId,
        });
      } catch (err) {
        this.apiError = true;
        return;
      }
      this.apiError = false;
      this.message = "正常に削除できました。"
      this.getTodo();
    },
    async createTask(todo) {
      let result;
      try {
        result = await api.post("/todo/create", {
          title: todo.title,
          note: todo.note,
          due_date: todo.due_date,
          is_completed: todo.is_completed,
        });
      } catch (err) {
        this.apiError = true;
        return;
      }
      this.apiError = false;
      this.message = "正常に保存できました。"
      this.getTodo()
      console.log(result);
    },
  },
};
</script>

<style lang="scss">
</style>