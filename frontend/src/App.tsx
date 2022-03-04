import React from 'react';
import './App.css';
import { Header } from './components/Header';
import { TodoList } from './components/TodoList';

function App() {
  return (
    <div className="App">
      <Header />
      <TodoList todos={[
        {title: "Grocery Shopping", isCompleted:true},
        {title: "Study for exam", isCompleted:false}
      ]} />
    </div>
  );
}

export default App;
