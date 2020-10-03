<template>
  <div class="container">
    <div class="row">
      <div class="col-md-12">
        <b-alert v-if="apiError" class="my-2" variant="danger" show dismissible>
          サーバーと通信ができません。しばらく時間をおいてからもう一度お試しください。
        </b-alert>
        <b-alert
          v-if="message"
          class="my-2"
          role="alert"
          variant="success"
          show
          dismissible
        >
          {{ message }}
        </b-alert>
      </div>
      <div class="col-md-4">
        <Form
          ref="form"
          :editTodoId="editTodoId"
          @createTask="createTask"
          @updateTodo="updateTodo"
          @updateEditTodoId="updateEditTodoId"
        />
      </div>
      <div class="col-md-8">
        <List
          :todos="todos"
          :apiError="apiError"
          :editTodoId="editTodoId"
          @callSetTodo="callSetTodo"
          @deleteTodo="deleteTodo"
          @getTodo="getTodo"
          @completeTodo="completeTodo"
          @updateEditTodoId="updateEditTodoId"
        />
      </div>
    </div>
  </div>
</template>

<script>
import api from "@/api";

import Form from "@/components/Form";
import List from "@/components/List";
export default {
  name: "todoApp",
  components: {
    Form,
    List,
  },
  data() {
    return {
      todos: [],
      apiError: false,
      message: "",
      editTodoId: 0,
    };
  },
  methods: {
    async getTodo() {
      let result;
      try {
        result = await api.get("/todo/get");
      } catch (err) {
        this.apiError = true;
        return;
      }
      this.todos = result.data;
      this.apiError = false;
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
      this.message = "正常に更新できました。";
      this.getTodo();
    },
    async updateTodo(todo) {
      let result;
      try {
        result = await api.post("/todo/update", {
          id: todo.id,
          title: todo.title,
          note: todo.note,
          due_date: todo.due_date,
          is_completed: false,
        });
      } catch (err) {
        this.apiError = true;
        return;
      }
      this.apiError = false;
      this.message = "正常に更新できました。";
      this.getTodo();
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
      this.message = "正常に削除できました。";
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
      this.message = "正常に保存できました。";
      this.getTodo();
    },
    callSetTodo(todo) {
      this.$refs.form.setTodo(todo);
    },
    updateEditTodoId(editTodoId) {
      this.editTodoId = editTodoId;
    },
  },
};
</script>
