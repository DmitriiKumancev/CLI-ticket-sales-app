# Ticket Purchase Simulation study cli-app

This Go application simulates the ticket purchase process for the "Go Conference." It allows users to book tickets and receive confirmation emails.

## Table of Contents

- [Overview](#overview)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)


## Overview

This application is a simple Go program that models the process of booking tickets for a conference. It maintains a list of booked tickets, verifies user input, and simulates sending confirmation emails.

## Prerequisites

To run this application, you need to have Go (Golang) installed on your system.

## Installation

1. Clone the repository:

   ```bash
   git clone <repository_url>

2. Navigate to the application directory:

   ```bash
   cd ticket-purchase-simulation

3. Build the application:

   ```bash
   go build main.go


## Usage

To use the application, follow these steps:

1. Run the application:

    ```bash
    ./main

* Follow the prompts to enter your first name, last name, email address, and the number of tickets you want to purchase.

* If your input is valid, the application will book the tickets, and you will receive a confirmation email simulation after a 10-second delay.

* You can check the remaining tickets and the list of booked tickets.

* The program will terminate when all tickets are booked.