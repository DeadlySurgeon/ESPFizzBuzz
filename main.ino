#include <WiFi.h>
#include <HTTPClient.h>

const char *ssid = "";
const char *password = "";

// Your Domain name with URL path or IP address with path
String serverName = "http://192.168.1.50:9000/fizzbuzz";
unsigned long timerDelay = 5000;
unsigned long lastTime = 0;

// the setup function runs once when you press reset or power the board
void setup()
{
  Serial.begin(115200);

  WiFi.begin(ssid, password);
  Serial.println("Connecting");
  while (WiFi.status() != WL_CONNECTED)
  {
    delay(500);
    Serial.print(".");
  }
  Serial.println("");
  Serial.print("Connected to WiFi network with IP Address: ");
  Serial.println(WiFi.localIP());
}

// the loop function runs over and over again forever
void loop()
{
  // Send an HTTP POST request every 10 minutes
  if ((millis() - lastTime) > timerDelay)
  {
    // Check WiFi connection status
    if (WiFi.status() == WL_CONNECTED)
    {
      play();
    }
    else
    {
      Serial.println("WiFi Disconnected");
    }
    lastTime = millis();
  }
}

int count = 0;

void play()
{
  if (count == 0)
  {
    Serial.println(start());
    count += 1;
    return;
  }
  String next = generate();
  Serial.println(next);

  String resp = submit(next);
  if (resp == "Error")
  {
    Serial.println("Got an error submitting?");
    return;
  }
  count++;
  Serial.println(resp);
}

String submit(String input)
{
  HTTPClient http;
  String serverPath = serverName + "?id=1&cmd=submit&entry=" + input;
  http.begin(serverPath.c_str());
  int respCode = http.GET();

  // All Good.
  if (respCode == 200)
  {
    return http.getString();
  }

  http.end();

  return "Error";
}

String start()
{
  HTTPClient http;
  String serverPath = serverName + "?id=1&cmd=new";
  http.begin(serverPath.c_str());

  http.GET();
  String resp = http.getString();

  http.end();

  return resp;
}

String generate()
{
  count++;
  return fizzbuzz(count);
}

// Generates a new FizzBuzz from int i
String fizzbuzz(int i)
{
  String s = String(i);
  bool set = false;

  if (i % 3 == 0)
  {
    s = "fizz";
    set = true;
  }
  if (i % 5 == 0)
  {
    if (!set)
    {
      s = "";
    }
    s += "buzz";
  }

  return s;
}
