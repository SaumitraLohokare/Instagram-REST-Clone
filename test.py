import requests
import datetime

def runTests():
    resp = requests.get("http://localhost:5000/users/1")
    assert resp.status_code == 200

    resp = requests.get("http://localhost:5000/posts/1")
    assert resp.status_code == 200

    resp = requests.get("http://localhost:5000/posts/2")
    assert resp.status_code == 200

    resp = requests.get("http://localhost:5000/posts/users/1/1")
    assert resp.status_code == 200

    print("All tests passed")

def addUser(id, name, email, password):
    resp = requests.post("http://localhost:5000/users", {"Id": id, "Name": name, "Email": email, "Password": password})
    if resp.status_code == 201:
        print("User created")
    else:
        print("Error creating user")

def addPost(id, userId, caption, imageURL):
    resp = requests.post("http://localhost:5000/users", {"Id": id, "UserId": userId, "Caption": caption, "ImageURL": imageURL, "Timestamp": datetime.datetime.now()})
    if resp.status_code == 201:
        print("Post created")
    else:
        print("Error creating post")
