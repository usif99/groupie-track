#Groupie Trackers#
Groupie Trackers consists on receiving a given API and manipulate the data contained in it, in order to create a site, displaying the information.

It will be given an API, that consists in four parts:

The first one, artists, containing information about some bands and artists like their name(s), image, in which year they began their activity, the date of their first album and the members.

The second one, locations, consists in their last and/or upcoming concert locations.

The third one, dates, consists in their last and/or upcoming concert dates.

And the last one, relation, does the link between all the other parts, artists, dates and locations.

Given all this you should build a user friendly website where you can display the bands info through several data visualizations (examples : blocks, cards, tables, list, pages, graphics, etc). It is up to you to decide how you will display it.

#Usage and how to run:
Clone the Repository:

git clone https://learn.reboot01.com/git/hnooh/groupie-tracker.git

Run the program:

go run main.go

OR

go run .

The server will start on http://localhost:8080

Open your web browser and navigate to http://localhost:8080 to access the page.

#Authors#
yousif maidan (ymaidan)
tayma fakhar (tfakhar)
maryam zaman (mzaman)


#explantion of the code:

*main.go*



------func openbrowser--->
1-The function takes a single argument, url, which is the URL to be opened in the web browser.

2-It initializes two variables, cmd and args, which will be used to construct the command to open the web browser.

3-The switch statement determines the appropriate command and arguments based on the user's operating system:
On Windows, the command is "cmd" and the arguments are "/c", "start", which will open the default web browser.
On macOS (Darwin), the command is "open", which will also open the default web browser.
For other operating systems (e.g., Linux, FreeBSD, OpenBSD, NetBSD), the command is "xdg-open", which is a common command for opening files and URLs on Linux systems.

4-The args slice is then appended with the url argument.

5-Finally, the exec.Command(cmd, args...).Start() function is called to execute the command and open the web browser with the specified URL.



-------func index handler---->

1-The function takes two arguments: w http.ResponseWriter and r *http.Request. These are the standard arguments for an HTTP request handler in Go.

2-It first checks if the request method is "GET". If it's not, it returns a 400 Bad Request error.

3-If the request method is "GET", it checks if the request URL path is not the root ("/"). If it's not, it renders the "error.html" template with a 404 Not Found status.

4-If the request URL path is the root ("/"), it parses the "index.html" template. If there's an error parsing the template, it returns a 500 Internal Server Error.

5-It then calls the create.GetArtists() function to retrieve a list of artists. If there's an error getting the artists, it returns a 500 Internal Server Error.

6-Finally, it executes the "index.html" template and passes the list of artists as the data to be rendered.



 ----artis handler--->

 Next, it extracts the artist name from the URL path. The artist name is expected to be part of the URL path, starting from the 8th character (e.g., /artist/john-doe).
It then iterates through the list of artists to find the one that matches the extracted artist name from the URL path.
If a matching artist is found, it executes the "profile.html" template and passes the artist data to be rendered.



*createDate.go*

purpose:

The purpose of this code is to provide a more user-friendly representation of the data retrieved from the API endpoints. It handles the data processing and formatting to make the information more accessible and understandable for the end-user.

Overall, this code demonstrates the use of Go to interact with RESTful APIs, parse JSON data, and transform the retrieved information into a structured format that can be easily consumed by other parts of the application.

Structs:

Artist: Defines the structure of an artist, including its ID, image, name, members, creation date, first album, and relations.
Relations: Defines the structure of the relations data, including the ID and a map of dates and locations.

Functions:

getRelations(): Retrieves the relations data from the /api/relation endpoint, parses the JSON response, and modifies the date and key formats to be more readable.
GetArtists(): Retrieves the artist data from the /api/artists endpoint, parses the JSON response, and combines the artist information with the corresponding relations data.
createBody(link string): Sends an HTTP GET request to the provided URL and returns the response body as a byte slice.