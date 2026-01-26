use std::io::{Read, Write};
use std::net::{TcpListener, TcpStream};

pub fn start_server() {
    let listener = TcpListener::bind("127.0.0.1:7890").unwrap();
    println!("Server started on port 7890");

    for stream in listener.incoming() {
        let stream = stream.unwrap();
        handle_client(stream);
    }
}

fn handle_client(mut stream: TcpStream) {
    // The buffer acts as a safety net
    // The OS will catch the bytes sent by the client and store them in the buffer until the code is ready to read them
    let mut buffer = [0; 512];

    loop {
        match stream.read(&mut buffer) {
            Ok(0) => {
                println!("Client disconnected");
                break;
            }
            Ok(n) => {
                let received_data = &buffer[..n];
                //checks for a null byte to determine the end of the message
                if received_data.contains(&0) {
                    println!("Message Ended");

                    let response = "Message Received\0";
                    stream.write_all(response.as_bytes()).unwrap();
                }
            }
            Err(e) => {
                println!("Error: {}", e);
                break;
            }
        }
    }
}
