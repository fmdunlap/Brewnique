import { useState } from "react";

function App() {
  const [id, setId] = useState<number>(0);
  const [loading, setLoading] = useState<boolean>(false);
  const [recipe, setRecipe] = useState<string>("");

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
      console.log(e.target.value);
      if (e.target.value === "") {
          return;
      }
      setId(parseInt(e.target.value));
  }

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
      e.preventDefault();
      setLoading(true);
      const response = await fetch(`http://localhost:8080/v1/recipe/${id}`);
      const data = await response.json();
      console.log(data);
      setLoading(false);
      setId(0);
      setRecipe(JSON.stringify(data, null, 2));
  }

  return (
    <>
        <h1>Brewnique</h1>
        <form onSubmit={handleSubmit}>
            <div className="flex flex-col">
                <label htmlFor="id">recipe id</label>
                <input type="number" id="id" name="id" value={id} onChange={handleChange} />
            </div>
            <button type="submit">
                {loading ? <span>Loading...</span> : <span>Submit</span>}
            </button>
        </form>
        <div className="flex flex-col">
            <label htmlFor="recipe">recipe</label>
            {recipe && <pre>{recipe}</pre>}
        </div>
    </>
  );
}

export default App;
