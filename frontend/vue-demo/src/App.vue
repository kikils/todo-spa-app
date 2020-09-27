<template>
  <div id="app">
    <Header />
    <div class="container">
      <div class="row">
        <div class="col-12">
          <br />
        </div>
        <div class="col-4">
          <Form @createTask="createTask" />
        </div>
        <div class="col-8">
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
      this.getTodo()
      console.log(result);
    },
  },
};
</script>

<style lang="scss">
</style>