      
Backend system built in Go for saving weather registries plus an alerting system
      
    Abstract

Starting from the problem we had, it has been build a system that exposes API endpoints to save weather registries. You can also get the list of the recorded registries. The system will hit the public wheather api endpoint to fetch the weather information. 

The solution did not need to expose the API, and could have run a function over the three cities asked. Knowing that, I decided to took the API-first approach.

Based on our requirements, an alert will be triggered if a bad weather state is created. To establish those states, both parameters limits and mathematical operations have been set in order to calculate the final weather state: temperature, humidity, wind speed and percentage of clouds in the sky (related to the presence of sunlight). 
There are 4 possible weather states: Bad, Neutral, Good and Undefined Weather. 

In future versions, the app could have ways to modify these weather parameters limits, store them in a repository, or be loaded from config files. For now, they are hardcoded in the program. Also the math operations could be re-thought to be more complex and accurated. 

For saving weather registries, here you have the endpoints. You might change {city_name}, to save a weather registry for the desired city. It will take the data, compute and save the result. 
Hit the get endpoint to retrieve what it has been stored. 

POST   localhost:5000/weather/{city_name}
GET    localhost:5000/weather

During the creating process, a notification will be triggered if required based on weather state, previously to the saving operation. If that fails, it will not reach the saving point.  

Weather service can be configured with different repos (memory and mysql) and notification systems (mock and twilio). Mock notification just logs to std.out that an alert has been triggered.

Twilio service is a functional notification system built with Twilio. In this case, my personal mobile phone has been submitted to Twilio as the receptor of the alerts. 

----------------------------------------------------------------------

Regarding architecture and development, DDD principles have been followed, taking into account that Go packages, interfaces, etc. might allow us to create a slighlty different folder structure than other languages.

Docker compose will launch the complete system(mysql repo and twilio service). Phpadmin is also launched to interact with the db easily with the UI.  

Build and run the system:

    docker compose up --build

