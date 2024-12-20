CC=gcc
CFLAGS=-Wall -g
SRC_DIR=src
OBJ_DIR=obj
BIN_DIR=bin
TARGET=$(BIN_DIR)/crush

# 文件路径
SRC=$(wildcard $(SRC_DIR)/*.c)
OBJ=$(SRC:$(SRC_DIR)/%.c=$(OBJ_DIR)/%.o)

all: $(TARGET)

$(TARGET): $(OBJ)
	$(CC) $(CFLAGS) -o $(TARGET) $(OBJ)

$(OBJ_DIR)/%.o: $(SRC_DIR)/%.c
	$(CC) $(CFLAGS) -c $< -o $@

clean:
	rm -rf $(OBJ_DIR)/*.o $(BIN_DIR)/crush

run: $(TARGET)
	./$(TARGET)
