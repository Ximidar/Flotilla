
String readline;
bool finline;


void setup(){
  Serial.begin(115200);  
  Serial.println("Starting!");
  
}
void loop() {

	if( Serial.available() ){
		while(Serial.available() > 0){
			char c = Serial.read();
			readline += c;
			if (c == '\n'){
				finline= true;
			}
		}

		if( finline ){
			Serial.println(readline);
			readline = "";
			finline = false;
		}
		
	}

	//Serial.println("Hello!");

}