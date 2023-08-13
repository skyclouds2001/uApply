# Store --- Local Storage

## What is this?
 
> This is a encapsulation for storage.
> 
> Defferent from VUEX, localStroage and sessionStorage are the way provided by brower.


## Why do this?

- To standard the api call which can avoid different developer use the bottom api in a way out of order .
- To better handle the storage need .


## What we do?

> We encapsulate two api of the brower : localStroage and sessionStorage

If you want to know the usage of them , you should go to check the information youself.


## How can I use them ?

If you want to use them in your component,you should follow:

**Use localStroage:**
```
this.localData("way","name",data);
```

**Use sessionStroage:**
```
this.sessionData("way","name",data);
```

**Props**
| props name | type | explain |
|------------|------|---------|
|    way     | string | select from ["get","set","clean"] |
|    name    | string | symbol of your data in the storage |
|    data    |  any   | data you store in the storage |
