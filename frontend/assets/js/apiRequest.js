
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
            console.log("ResponseJs", jsResponse);
            return jsResponse;
        } catch (error) {
            console.log("Response is not valid JSON:", textResponse);
            console.log("END:");
            throw error;
        }
        
    }catch(error){
        console.log(`Error Fetch data from http://localhost:8080/${url}`, error)
    }
}