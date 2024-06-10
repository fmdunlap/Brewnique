import { useState } from "react";

function NewRecipeForm() {
    const [name, setName] = useState<string>("");
    const [ingredients, setIngredients] = useState<string[]>([]);
    const [instructions, setInstructions] = useState<string[]>([]);
    const [loading, setLoading] = useState<boolean>(false);

    const handleNameChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        setName(e.target.value);
    }

    const handleIngredientsChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        setIngredients(e.target.value.split(","));
    }

    const handleInstructionsChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        setInstructions(e.target.value.split(","));
    }

    const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        setLoading(true);
        const response = await fetch(`http://localhost:8080/v1/recipe`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({
                name: name,
                ingredients: ingredients,
                instructions: instructions,
            }),
        });
        const data = await response.json();
        console.log(data);
        setLoading(false);
    }

    return (
        <form onSubmit={handleSubmit} className="flex flex-col gap-2">
            <div className="flex flex-row gap-2">
                <label className="font-bold my-auto" htmlFor="name">name</label>
                <input className="flex-1 bg-gray-50 rounded-lg border-2 border-gray-300 p-2" type="text" id="name" name="name" value={name} onChange={handleNameChange} />
            </div>
            <div className="flex flex-row gap-2">
                <label className="font-bold my-auto" htmlFor="ingredients">ingredients</label>
                <input className="flex-1 bg-gray-50 rounded-lg border-2 border-gray-300 p-2" type="text" id="ingredients" name="ingredients" value={ingredients} onChange={handleIngredientsChange} />
            </div>
            <div className="flex flex-row gap-2">
                <label className="font-bold my-auto" htmlFor="instructions">instructions</label>
                <input className="flex-1 bg-gray-50 rounded-lg border-2 border-gray-300 p-2" type="text" id="instructions" name="instructions" value={instructions} onChange={handleInstructionsChange} />
            </div>
            <button type="submit" className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded">
                {loading ? <span>Loading...</span> : <span>Submit</span>}
            </button>
        </form>
    )
}

function RecipeList() {
    const [recipes, setRecipes] = useState<any[]>([]);
    const [extraData, setExtraData] = useState<any>({});

    const handleListRecipes = async () => {
        const response = await fetch(`http://localhost:8080/v1/recipes`);
        const data = await response.json();
        console.log(data);
        setRecipes(data);
    }

    const handleDeleteRecipe = async (id: number) => {
        const response = await fetch(`http://localhost:8080/v1/recipe/${id}`, {
            method: "DELETE",
        });
        const data = await response.json();
        console.log(data);
        if (data.status === "error") {
            alert("Error deleting recipe");
        } else {
            handleListRecipes();
        }
    }

    const handleGetRecipe = async (id: number) => {
        const response = await fetch(`http://localhost:8080/v1/recipe/${id}`);
        const data = await response.json();
        console.log(data);
        setExtraData(data);
    }

    return (
        <>
            <div className="flex flex-row gap-2">
                <h2 className="text-2xl font-bold">List Recipes</h2>
                <button onClick={handleListRecipes} className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded">Refresh</button>
            </div>
            <div className="p-4 bg-amber-50 rounded-lg">
                {recipes.map((recipe) => (
                    <div className="flex flex-col gap-2">
                        <p className="font-bold">Name: {recipe.name}</p>
                        <p className="font-bold">Ingredients: {recipe.ingredients.join(", ")}</p>
                        <p className="font-bold">Instructions: {recipe.instructions.join(", ")}</p>
                        <button onClick={() => handleDeleteRecipe(recipe.id)} className="bg-red-500 hover:bg-red-700 text-white font-bold py-2 px-4 rounded">Delete</button>
                        <button onClick={() => handleGetRecipe(recipe.id)} className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded">Get</button>
                    </div>
                ))}
            </div>
            <div className="p-4 bg-amber-100 rounded-lg">
                {extraData && <pre>{JSON.stringify(extraData, null, 2)}</pre>}
            </div>
        </>
    )

}


function App() {

    return (
        <div className="flex flex-col gap-2 p-4">
            <h1 className="text-3xl font-bold">Brewnique</h1>
            <h2 className="text-2xl font-bold">New Recipe</h2>
            <NewRecipeForm/>
            <RecipeList/>
        </div>
    );
}

export default App;
