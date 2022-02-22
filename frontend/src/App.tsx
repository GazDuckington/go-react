import React, { useEffect, useState } from 'react'
import './App.css'

function App() {
  const url1 = 'http://localhost:5000/api/books';
  const [books, setBooks] = useState([] as any[]);
  const [loading, setLoading] = useState(true);

  async function fetchBooks(){
    const res = await fetch(url1, {method: 'GET'});
    const data = await res.json();    
    if (res.ok) {
      console.log(data);
      setBooks(data)
      setLoading(false)
    }
    else {
      throw new Error(data);      
    }
  }

  useEffect(() => {
    fetchBooks();
  }, [])

  return (
    <div className="App">
      <title>Home</title>

      {loading ? <p>loading...</p> : <p>books...</p> }

      {books.map((item, i) => {
          return <li key={i}>{item.title}</li>
        })}
    </div>
  )
}

export default App
