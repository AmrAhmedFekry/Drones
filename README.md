This solution based on my basic MVC file structure [Sakara](https://github.com/AmrAhmedFekry/Sakara).


# System Design  

![ER Diagram](ER.png)

As Shown in diagram we have collapse the medications items in drone_load that gave us the ability to track the load, and make logs for it.



# API DOCS




# Setting API DOCS

# Base URL
http://localhost:8080

# Other resources 

 
# Headers

Locale: ar|en 

Accept : application/json

# API 

| Route                        | Request Method | Parameters | Response  |
| -----------                  | -----------    |----------- |---------- |
| /api/drone/register          | POST           |  [Register Parameters](#RegisterDrone)|[Response](#Response)|
|/api/medication               |POST           |  [Register Parameters](#RegisterMedication)|[Response](#Response)     |
| /api/load          | POST           |  [Load Parameters](#LoadDrone)|[Response](#Response)|

| /api/upload          | POST           |  [Upload Parameters](#UploadMedia)|[Response](#UploadResponse)|

| /available/loading/drones          | GET  |  [-](#AvailableResponse)|[Response](#AvailableResponse)|


| /drone/battery_level/{droneId}          | GET  |  [-](#showBatteryResponse)|[Response](#showBatteryResponse)|

| /change_drone_load_state          | PUT  |  [-](#ChangeDroneLoadRequest)|[Response](#ChangeDroneLoadResponse)|

# <a name="UploadMedia"> </a> Upload Media 

```json
{
    "image" : "File",
} 
```


# <a name="UploadResponse"> </a> Upload Media 

```json
{
    "fileName": "string",
    "filePath": "String",
    "message": "Your file has been successfully uploaded."
} 
```




# <a name="RegisterDrone"> </a> Register new Drone 

```json
{
    "serial_number" : "String",
    "model" : "String",
    "weight_limit" : "int"
} 
```

# <a name="LoadDrone"> </a> Load Drone with medications 

```json
{
    "drone_id" : "Int",
    "medication_items" : "ArrayOfIds",
    "latitude" : "int", 
    "longitude" : "int"
}
```


# <a name="RegisterMedication"> </a> Register Medication

```json
{
    "code" : "String",
    "name" : "String",
    "weight" : "int", 
    "image" : "String"
    } 
```

`Note` You can Use the upload API, to upload medication image and use the returned fileName from response as image param.


# <a name="Response"> </a> Responses 

## Validation error 
__*Response code : 401*__

```json 
{
    "code": 401,
    "errors": [],
    "message": "Validation Error",
    "payload": null
}
```
## Success  
__*Response code : 201*__
```json 
{
    "code": 200,
    "errors": null,
    "message": "Success",
    "payload": [
        {
        }
    ]
}
```

`Note` status code will 201 if the Request is POST




# <a name="AvailableResponse"> </a> Available Drones for loading Response 

```json
{
    "code": 200,
    "errors": null,
    "message": "Success",
    "payload": [
        {
            "BatteryCapacity": "Int",
            "DroneModel": "String",
            "DroneState": "String",
            "SerialNumber": "String",
            "WeightLimit": "int",
            "id": "Int"
        }
    ]
}} 
```


# <a name="showBatteryResponse"> </a> Show Battery level Response 

```json
{
    "code": 200,
    "errors": null,
    "message": "Success",
    "payload": 100
}}}
```

# <a name="ChangeDroneLoadRequest"> </a> Change Drone Load Request 

```json
{
    "drone_load_id" : 24,
    "drone_load_state" : "LOADED"
}}} 

```
# <a name="ChangeDroneLoadResponse"> </a> Change Drone Load Response 

```json
{
    "code": 200,
    "errors": null,
    "message": "Success",
    "payload": "Drone Load State Changed"
}}} 

