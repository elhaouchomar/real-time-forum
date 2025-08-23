
export async function apiRequest(url){
    try
    {
        const response = await fetch(`/${url}`, {
            method: "POST",
        })
        const textResponse = await response.text();
        try {
            const jsResponse = JSON.parse(textResponse);

            return jsResponse;
        } catch (error) {

            throw error;
        }
        
    }catch(error){
        console.log(`Error Fetch data from http://localhost:9090/${url}`, error)
    }
}