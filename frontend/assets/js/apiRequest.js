
export async function apiRequest(url){
    console.log(`Api Request Made to = ${url}`);
    try
    {
        const response = await fetch(`/${url}`, {
            method: "POST",
        })
        const textResponse = await response.text();
        try {
            const jsResponse = JSON.parse(textResponse);
            console.log("Result of jsResponse ", jsResponse);
            return jsResponse;
        } catch (error) {
            console.log("Response is not valid JSON:", textResponse);
            throw error;
        }
        
    }catch(error){
        console.log(`Error Fetch data from http://localhost:9090/${url}`, error)
    }
}