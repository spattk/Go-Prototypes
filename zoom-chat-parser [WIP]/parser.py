import re

def append_to_store(store, msg):
    store.append(msg)

def get_parent_msg(store, msg):
    cleaned_text = msg.replace("...", "")
    print("CLEANED_TEXT", cleaned_text)
    for idx,m in enumerate(store):
        print("M", m)
        if cleaned_text in m:
            return idx
        
    return -1
            

def is_parent_msg(msg):
    if "Reacted to" in msg or "Replying to" in msg:
        return False
    return True

def extract_messages_between_timestamps(chat_text):
    # Regex to capture messages between timestamps
    message_pattern = re.compile(r'\d{2}:\d{2}:\d{2} From .*? to .*?:[\s\S]*?(?=\d{2}:\d{2}:\d{2} From|$)')
    store = []
    
    # Find all messages between timestamps
    messages = message_pattern.findall(chat_text)
    for message in messages:
        print("---------------")
        is_parent = is_parent_msg(message)
        
        if is_parent:
            parent_idx = None
            append_to_store(store, message)
        else:
            parent_idx = get_parent_msg(store, message)
            # append_to_store_replies(store, message, parent_idx)

        print(is_parent, parent_idx, message)
    # return messages

def read_chat_from_file(file_path):
    with open(file_path, 'r', encoding='utf-8') as file:
        chat_text = file.read()
    return chat_text

# Replace 'zoom_chat.txt' with the path to your chat file
file_path = 'min_chat.txt'
chat_text = read_chat_from_file(file_path)

extract_messages_between_timestamps(chat_text)